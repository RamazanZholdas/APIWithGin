package main

import (
	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/initialData"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = "8080"
)

func init() {
	file := ginLogs.SetupLogOutput()
	defer file.Close()

	initialData.PutData()
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLogs.Logger())

	r.GET("/getSongs", routes.GetSongs)
	r.GET("/getSongById/:id", routes.GetSongById)
	r.POST("/createSong", routes.CreateSong)
	r.PUT("/updateSong/:id", routes.UpdateSong)
	r.DELETE("/deleteSong/:id", routes.DeleteSong)

	r.Run(":" + port)
}
