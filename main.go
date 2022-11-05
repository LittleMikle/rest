package main

import (
	"github.com/LittleMikle/rest.git/controllers"
	"github.com/LittleMikle/rest.git/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	r := gin.Default()

	database.Init()
	//models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	//r.GET("/books/:id", controllers.FindBook)
	//r.PATCH("/books/:id", controllers.UpdateBook)
	//r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
