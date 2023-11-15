package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/miniapi_version2/database"
	"github.com/teerapoom/miniapi_version2/database/model"
)

func DeleteProduct(c *gin.Context) {
	var Movie model.Movie
	id := c.Param("id") // รับค่าพารามิเตอร์ "id" จาก URL และเก็บไว้ในตัวแปร id.

	// ตรวจสอบว่ามีสินค้าที่ต้องการลบหรือไม่
	if err := database.Db.Preload("Director").First(&Movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return //Stop
	}
	// ลบสินค้าออกจากฐานข้อมูล
	if err := database.Db.Delete(&Movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete Movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Ok",
		"message": "Movie deleted successfully"})
}
