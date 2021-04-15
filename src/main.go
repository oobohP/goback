package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oobohP/goback/controllers"
	"github.com/oobohP/goback/models"
)

func main() {

	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/books", controllers.FindBooks)

	r.Run()
}
