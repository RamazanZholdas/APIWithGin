package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/middleware"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file:\n", err)
	}
	databaseConn.ConnectToDB()
	databaseConn.MigrateModelToDB()
}

func main() {
	file := ginLogs.SetupLogOutput()
	defer file.Close()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginLogs.Logger())
	r.Use(middleware.SetCorsMiddleware())

	r.GET("/getAllSongs", routes.GetAllSongs)
	r.GET("/getSongById/:id", routes.GetSongById)
	r.POST("/createSong", routes.CreateSong)
	r.PUT("/updateSong/:id", routes.UpdateSong)
	r.DELETE("/deleteSong/:id", routes.DeleteSong)

	fmt.Println("Running on port:", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
