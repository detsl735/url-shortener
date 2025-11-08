package main

import (
	"fmt"

	"github.com/Detsl735/url-shortener/handler"
	"github.com/Detsl735/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Let's go to start!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error: %v", err))
	}
}
