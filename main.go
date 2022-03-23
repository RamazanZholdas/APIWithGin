package main

import (
	"time"

	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/RamazanZholdas/APIWithGin/structs"
	"github.com/gin-gonic/gin"
)

var (
	songs = []structs.Song{}
)

const (
	port = "8080"
)

func init() {
	file := ginLogs.SetupLogOutput()
	defer file.Close()

	songs = append(songs, structs.Song{
		Name:     "Bohemian Rhapsody",
		Duration: time.Now(),
		Genre:    "R&B",
		Artist: &structs.Artist{
			FirstName: "Freddie",
			LastName:  "Mercury",
			Label:     "Queen",
		},
	})
	songs = append(songs, structs.Song{
		Name:     "Crazy Train",
		Duration: time.Now().Add(time.Hour * 3),
		Genre:    "Metal",
		Artist: &structs.Artist{
			FirstName: "Ozzy",
			LastName:  "Osbourne",
			Label:     "Black Sabbath",
		},
	})
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
