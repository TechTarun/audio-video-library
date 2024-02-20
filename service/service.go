package service

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"example.com/avlib/db"
)

type Album struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price int `json:"price"`
}

func GetAllAlbums(c *gin.Context) {
	dbconn := db.Connect()
	rows, err := dbconn.Query("SELECT * FROM album")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Fatal(err)
		}

		albums = append(albums, album)
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Query("id")
	dbconn := db.Connect()

	var album Album
	dbconn.QueryRow("SELECT * FROM album WHERE id= $1", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	c.JSON(http.StatusOK, album)
}

func GetAlbumsByArtist(c *gin.Context) {
	artist := c.Query("artist")
	dbconn := db.Connect()
	rows, err := dbconn.Query("SELECT * FROM album WHERE artist = $1", artist)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Fatal(err)
		}

		albums = append(albums, album)
	}

	c.JSON(http.StatusOK, albums)
}

func AddNewAlbum(c *gin.Context) {
	dbconn := db.Connect()
	var newalbum Album
	if err := c.BindJSON(&newalbum); err != nil {
		log.Fatal(err)
	}

	var newid int
	err := dbconn.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", newalbum.Title, newalbum.Artist, newalbum.Price).Scan(&newid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newid)

	c.JSON(http.StatusCreated, newid)
}

func RemoveAlbumByID(c *gin.Context) {
	dbconn := db.Connect()
	id_to_remove := c.Query("id")
	err := dbconn.QueryRow("SELECT id FROM album WHERE id = $1", id_to_remove).Scan(&id_to_remove)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbconn.Exec("DELETE FROM album WHERE id = $1", id_to_remove)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, id_to_remove)
}
