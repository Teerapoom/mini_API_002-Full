package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/miniapi_version2/database"
	"github.com/teerapoom/miniapi_version2/database/model"
)

type Movie struct {
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Score    float64   `json:"score"`
	Director *Director `json:"Director"`
}

type Director struct {
	Fistname string `json:"fistname"`
	Lastname string `json:"lastname"`
	// Director *Director
}

var ListMovies []Movie

func PostMovie(c *gin.Context) {
	var NewMovie model.Movie
	if err := c.BindJSON(&NewMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var moviesExist model.Movie
	database.Db.Where("title = ?", NewMovie.Title).First(&moviesExist)
	if moviesExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Movie Exists"})
		return
	}

	// เพิ่มรายการหนังใหม่เข้าไปใน slice
	// database.Db.Create(&NewMovie)
	result := database.Db.Create(&NewMovie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"moive": NewMovie})
}
