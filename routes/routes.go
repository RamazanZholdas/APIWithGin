package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllSongs(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var result []structs.Song

	databaseConn.Db.Find(&result)
	json.NewEncoder(c.Writer).Encode(result)
}

func GetSongById(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Param("id")

	var song structs.Song

	query := databaseConn.Db.First(&song, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	} else {
		json.NewEncoder(c.Writer).Encode(song)
	}
}

func CreateSong(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var song structs.Song

	err := json.NewDecoder(c.Request.Body).Decode(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json": errors.New("cant decode")})
		return
	}

	databaseConn.Db.Select("Name", "Duration", "Genre", "Artist").Create(&song)
	json.NewEncoder(c.Writer).Encode(song)
}

func UpdateSong(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Param("id")

	var song structs.Song
	var updatedSong structs.SimpleSong

	query := databaseConn.Db.First(&song, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	} else {
		err := json.NewDecoder(c.Request.Body).Decode(&updatedSong)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Json": errors.New("cant decode")})
			return
		}

		song.Name = updatedSong.Name
		song.Artist = updatedSong.Artist
		song.Genre = updatedSong.Genre
		song.Duration = updatedSong.Duration
		databaseConn.Db.Save(&song)

		json.NewEncoder(c.Writer).Encode(updatedSong)
	}
}

func DeleteSong(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Param("id")

	var song structs.Song

	query := databaseConn.Db.First(&song, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	} else {
		databaseConn.Db.Delete(&song)

		json.NewEncoder(c.Writer).Encode("song deleted")
	}
}
