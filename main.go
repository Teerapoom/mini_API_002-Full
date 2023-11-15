package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/teerapoom/miniapi_version2/Controllers"
	"github.com/teerapoom/miniapi_version2/database"
)

func main() {
	r := gin.Default()
	database.InitDB()
	r.Use(cors.Default())
	r.GET("/Getbook", controllers.GetAllMoive)
	r.GET("/Getbook/:id", controllers.GetMoiveByID)
	r.POST("/addbook", controllers.PostMovie)
	r.PUT("/Updatebook/:id", controllers.UpdateMovie)
	r.DELETE("/Deletebook/:id", controllers.DeleteProduct)
	r.Run("localhost:3030") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
