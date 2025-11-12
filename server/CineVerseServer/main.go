package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "github.com/adityapandeydev/CineVerse/Server/CineVerseServer/controllers"
	"github.com/adityapandeydev/CineVerse/Server/CineVerseServer/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, CineVerse!")
	})

	var client *mongo.Client = database.Connect()
	router.GET("/movies", controller.GetMovies(client))

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}