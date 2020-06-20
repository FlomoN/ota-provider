package main

import (
	"github.com/flomon/ota-provider/src/db"
	"github.com/flomon/ota-provider/src/monitor"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	store := db.Load()
	monitor.Init(store)

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", false)))

	r.Use(static.Serve("/bin", static.LocalFile("./data", false)))

	//r.Static("/bin", "./data")

	r.GET("/data", func(c *gin.Context) {
		c.JSON(200, store)
	})

	r.POST("/force", func(c *gin.Context) {
		type ForceInfo struct {
			ID int `json:"id"`
		}
		requestBody := ForceInfo{}
		c.Bind(&requestBody)
		monitor.ForceDeviceUpdate(requestBody.ID)
		c.Status(200)
	})

	r.POST("/remove", func(c *gin.Context) {
		type DeleteInfo struct {
			ID int `json:"id"`
		}
		requestBody := DeleteInfo{}
		c.Bind(&requestBody)
		store.Remove(requestBody.ID)
		c.Status(200)
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
			MQTT string
		}

		requestBody := GithubCreds{}
		c.Bind(&requestBody)
		store.SetCreds(requestBody.Name, requestBody.Pass, requestBody.MQTT)
		c.Status(200)
	})

	r.Run(":3001")
}
