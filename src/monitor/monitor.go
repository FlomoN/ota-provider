package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/flomon/ota-provider/src/db"
)

type githubResponse struct {
	TagName string `json:"tag_name"`
	ID      int    `json:"id"`
	Assets  []struct {
		Link string `json:"browser_download_url"`
	} `json:"assets"`
}

var store *db.Store

// Init initializes the Monitor with a given Store and starts the watching coroutine
func Init(_store *db.Store) {
	store = _store

	ticker := time.NewTicker(60 * time.Second)
	go monitor(ticker.C)
}

func monitor(t <-chan time.Time) {

	for {
		if store.Init {
			for _, w := range store.Watch {
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
					fmt.Println(githubData)
				} else {
					fmt.Println(resp.Status)
				}

			}
		}
		<-t
	}
}
