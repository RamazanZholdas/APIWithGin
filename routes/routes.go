package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {

}

func GetSongById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id:", id)
}

func CreateSong(c *gin.Context) {

}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id:", id)
}

func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id:", id)
}
