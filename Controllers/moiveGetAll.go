package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/miniapi_version2/database"
	"github.com/teerapoom/miniapi_version2/database/model"
)

func GetAllMoive(c *gin.Context) {
	var Movie []model.Movie
	err := database.Db.Preload("Director").Find(&Movie).Error //Preload ใช่ตัวที่เชื่อม
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "Movie_List": Movie})

}

/*
คุณต้องการโหลดข้อมูลจากตาราง Director ที่เกี่ยวข้องกับแต่ละรายการในตาราง Movie.
*/
