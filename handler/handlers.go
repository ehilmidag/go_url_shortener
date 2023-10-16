package handler

import (
	"github.com/ehilmidag/go_url_shortener/shortener"
	"github.com/ehilmidag/go_url_shortener/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var createUrlRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&createUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := shortener.GenerateShortUrl(createUrlRequest.LongUrl, createUrlRequest.UserId)
	storage.SaveUrlMapping(shortUrl, createUrlRequest.LongUrl, createUrlRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
