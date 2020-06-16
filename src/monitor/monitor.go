package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/flomon/ota-provider/src/db"
)

type githubResponse struct {
	TagName string `json:"tag_name"`
	ID      int    `json:"id"`
	Assets  []struct {
		Link string `json:"url"`
	} `json:"assets"`
}

var store *db.Store
var mqttClient *MQTT.Client

// Init initializes the Monitor with a given Store and starts the watching coroutine
func Init(_store *db.Store) {
	store = _store
	mqttClient = createMQTTClient()

	ticker := time.NewTicker(60 * time.Second)
	go monitor(ticker.C)
}

func monitor(t <-chan time.Time) {

	for {
		if store.Init {
			for ix, w := range store.Watch {
				client := &http.Client{}
				fmt.Println("https://api.github.com/repos/" + store.GhName + "/" + w.Repo + "/releases")
				req, _ := http.NewRequest("GET", "https://api.github.com/repos/"+store.GhName+"/"+w.Repo+"/releases", nil)
				req.SetBasicAuth(store.GhName, store.GhToken)
				req.Header.Set("If-None-Match", w.ETag)

				resp, _err := client.Do(req)
				if _err != nil {
					log.Fatal("Error Reading Repo " + w.Repo)
				}
				defer resp.Body.Close()
				if resp.StatusCode == 200 {
					body, _ := ioutil.ReadAll(resp.Body)
					githubData := []githubResponse{}
					json.Unmarshal(body, &githubData)
					store.Watch[ix].UpdateEIV(store, resp.Header.Get("ETag"), githubData[0].ID, githubData[0].TagName)
					if githubData[0].ID != w.ReleaseID && len(githubData[0].Assets) > 0 {
						download, _ := http.NewRequest("GET", githubData[0].Assets[0].Link, nil)
						download.SetBasicAuth(store.GhName, store.GhToken)
						download.Header.Set("Accept", "application/octet-stream")
						downloadResp, _err := client.Do(download)
						if _err != nil {
							panic(_err)
						}
						defer downloadResp.Body.Close()

						downloadContent, _ := ioutil.ReadAll(downloadResp.Body)

						pwd, _ := os.Getwd()
						binary, _ := os.Create(filepath.Join(pwd, "./data/"+w.Repo+".png"))
						binary.Write(downloadContent)
						binary.Close()

						(*mqttClient).Publish("update/"+w.Device, 2, false, "")
					}
				} else {
					fmt.Println(resp.Status)
				}

			}
		}
		<-t
	}
}

func createMQTTClient() *MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + store.MQTTHost)
	opts.SetClientID("ota-updater")

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	c.Publish("ping", 0, false, "ota-updater")

	return &c
}
