package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/rahmanfadhil/gin-bookstore/controllers" // new
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)

	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
