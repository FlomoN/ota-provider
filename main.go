package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Watcher struct {
	Repo   string
	Device string
	Id     int
}

type Store struct {
	Init    bool
	GhName  string
	GhToken string
	Watch   []Watcher
}

var store Store

func main() {

	loadData()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hallo")
	})

	r.Run(":3000")
}

func loadData() {

	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	f, err := ioutil.ReadFile(filepath.Join(pwd, "./data/store.json"))
	if os.IsNotExist(err) {
		os.Mkdir(filepath.Join(pwd, "./data"), 0777)
		_f, _ := os.Create(filepath.Join(pwd, "./data/store.json"))
		store = Store{false, "", "", []Watcher{}}
		x, _ := json.Marshal(store)
		_f.Write(x)
		_f.Close()
	} else {
		fmt.Println(string(f))
		err = json.Unmarshal(f, &store)
	}

}
