package main

import(
	"github.com/gin-gonic/gin"
	"example.com/avlib/service"
)

func ginEngine() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", service.GetAllAlbums)
	router.GET("/album-by-id", service.GetAlbumByID)
	router.GET("/albums-by-artist", service.GetAlbumsByArtist)
	router.POST("/albums", service.AddNewAlbum)
	router.DELETE("/albums", service.RemoveAlbumByID)
	return router
}

func main() {
	router := ginEngine()
	router.Run("localhost:8080")
}