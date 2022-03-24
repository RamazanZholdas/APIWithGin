package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/RamazanZholdas/APIWithGin/initialData"
	"github.com/RamazanZholdas/APIWithGin/structs"
	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {
	if err := c.ShouldBindJSON(&initialData.Songs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.NewEncoder(c.Writer).Encode(initialData.Songs)
}

func GetSongById(c *gin.Context) {
	if err := c.ShouldBindJSON(&initialData.Songs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	didntFound := false

	for index := range initialData.Songs {
		if strconv.Itoa(initialData.Songs[index].Id) == id {
			json.NewEncoder(c.Writer).Encode(initialData.Songs[index])
			didntFound = true
			break
		}
	}
	if !didntFound {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	}
}

func CreateSong(c *gin.Context) {
	if err := c.ShouldBindJSON(&initialData.Songs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var song structs.Song

	err := json.NewDecoder(c.Request.Body).Decode(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
		return
	}

	initialData.Songs = append(initialData.Songs, song)
	json.NewEncoder(c.Writer).Encode(initialData.Songs)
}

func UpdateSong(c *gin.Context) {
	if err := c.ShouldBindJSON(&initialData.Songs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	didntFound := false

	for index := range initialData.Songs {
		if strconv.Itoa(initialData.Songs[index].Id) == id {
			initialData.Songs = remove(initialData.Songs, index)

			var song structs.Song

			err := json.NewDecoder(c.Request.Body).Decode(&song)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
				return
			}

			initialData.Songs = append(initialData.Songs, song)
			json.NewEncoder(c.Writer).Encode(initialData.Songs)

			didntFound = true
			break
		}
	}

	if !didntFound {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	}

}

func DeleteSong(c *gin.Context) {
	if err := c.ShouldBindJSON(&initialData.Songs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	didntFound := false

	for index := range initialData.Songs {
		if strconv.Itoa(initialData.Songs[index].Id) == id {
			initialData.Songs = remove(initialData.Songs, index)
			didntFound = true
			break
		}
	}
	if !didntFound {
		c.String(http.StatusBadRequest, "id does not exist")
	}

	json.NewEncoder(c.Writer).Encode(initialData.Songs)
}

func remove(slice []structs.Song, s int) []structs.Song {
	return append(slice[:s], slice[s+1:]...)
}
