package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "github.com/Amdadul-HQ/go_lang_app.git/go_app/server/movie_server/controllers"
)

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	// API routes
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello Movie Server")
	})

	router.GET("/movies",controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
