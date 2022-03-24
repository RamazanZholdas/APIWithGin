package main

import (
	"fmt"

	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/initialData"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/gin-gonic/gin"
)

const (
	port = "8080"
)

func init() {
	initialData.PutData()
}

func main() {
	file := ginLogs.SetupLogOutput()
	defer file.Close()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLogs.Logger())

	r.GET("/getSongs", routes.GetSongs)
	r.GET("/getSongById/:id", routes.GetSongById)
	r.POST("/createSong", routes.CreateSong)
	r.PUT("/updateSong/:id", routes.UpdateSong)
	r.DELETE("/deleteSong/:id", routes.DeleteSong)

	fmt.Println("Running on port:", port)
	r.Run(":" + port)
}
