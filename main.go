package main

import (
	"fmt"
	"github.com/ehilmidag/go_url_shortener/handler"
	"github.com/ehilmidag/go_url_shortener/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey go url shortener",
		})
	})

	r.POST("/create-short-url", func(context *gin.Context) {
		handler.CreateShortUrl(context)
	})
	r.GET("/:shortUrl", func(context *gin.Context) {
		handler.HandleShortUrlRedirect(context)
	})
	storage.InitializeStorage()
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("failed to start server - Error: %v", err))
	}
}
