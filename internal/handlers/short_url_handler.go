package handlers

import (
	"net/http"
	"url-shortener/internal/shortener"

	"github.com/gin-gonic/gin"
)

func UrlShortener(c *gin.Context) {
	var body struct {
		originURL string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty field",
		})
	}

	result := shortener.GenerateShortUrl(body.originURL)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your short url " + result,
	})
}
