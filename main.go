package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL shortener !",
		})
	})

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error: %v", err))
	}
}
