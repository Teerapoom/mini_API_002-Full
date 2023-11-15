package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/miniapi_version2/database"
	"github.com/teerapoom/miniapi_version2/database/model"
)

func UpdateMovie(c *gin.Context) {
	var updateInfo model.Movie
	id := c.Param("id") // Retrieve the "id" parameter from the URL

	// Find the movie by ID
	var movie model.Movie
	if err := database.Db.Preload("Director").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Read the data from the request body
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the movie in the database
	if err := database.Db.Model(&movie).Updates(updateInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Ok",
		"movie":  movie})
}
