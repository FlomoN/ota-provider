package monitor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flomon/ota-provider/src/db"
)

var store *db.Store

// Init initializes the Monitor with a given Store and starts the watching coroutine
func Init(_store *db.Store) {
	store = _store

	ticker := time.NewTicker(3 * time.Second)
	go monitor(ticker.C)
}

func monitor(t <-chan time.Time) {

	for {
		if store.Init {
			for _, w := range store.Watch {
				resp, _ := http.Get("http://localhost:3001/data")
				fmt.Println(resp.Body, w)
			}
		}
		<-t
	}
}
