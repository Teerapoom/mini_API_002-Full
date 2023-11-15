package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/miniapi_version2/database"
	"github.com/teerapoom/miniapi_version2/database/model"
)

func GetMoiveByID(c *gin.Context) {
	var moiveID model.Movie
	id := c.Param("id")

	//  เชื่อม -> First เป็นคำสั่งการค้นหา
	if err := database.Db.Preload("Director").First(&moiveID, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return //Stop
	}

	c.JSON(http.StatusOK, gin.H{"Movie": moiveID})
}
