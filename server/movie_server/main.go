package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	// API routes
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello Movie Server")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
