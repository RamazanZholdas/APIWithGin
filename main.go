package main

import (
	"fmt"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	port = "8080"
)

func init() {
	godotenv.Load()
	databaseConn.ConnectToDB()
	databaseConn.SyncDB()
}

func main() {
	file := ginLogs.SetupLogOutput()
	defer file.Close()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLogs.Logger())

	r.GET("/getAllSongs", routes.GetAllSongs)
	r.GET("/getSongById/:id", routes.GetSongById)
	r.POST("/createSong", routes.CreateSong)
	r.PUT("/updateSong/:id", routes.UpdateSong)
	r.DELETE("/deleteSong/:id", routes.DeleteSong)

	fmt.Println("Running on port:", port)
	r.Run(":" + port)
}
