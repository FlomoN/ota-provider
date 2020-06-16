package main

import (
	"github.com/flomon/ota-provider/src/db"
	"github.com/flomon/ota-provider/src/monitor"
	"github.com/gin-gonic/gin"
)

func main() {

	store := db.Load()
	monitor.Init(store)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hallo")
	})

	r.GET("/data", func(c *gin.Context) {
		c.JSON(200, store)
	})

	r.POST("/add", func(c *gin.Context) {
		type RepoEntry struct {
			Repo   string
			Device string
		}
		requestBody := RepoEntry{}
		c.Bind(&requestBody)

		store.Add(db.Watcher{Repo: requestBody.Repo, Device: requestBody.Device, ETag: "", Version: "", ReleaseID: 0})
		c.Status(200)
	})

	r.POST("/creds", func(c *gin.Context) {
		type GithubCreds struct {
			Name string
			Pass string
		}

		requestBody := GithubCreds{}
		c.Bind(&requestBody)
		store.SetCreds(requestBody.Name, requestBody.Pass)
		c.Status(200)
	})

	r.Run(":3001")
}
