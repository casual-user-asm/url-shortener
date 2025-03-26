package handlers

import (
	"net/http"
	"url-shortener/internal/shortener"

	"github.com/gin-gonic/gin"
)

var urls = make(map[string]string)

func GetOriginalURL(short string) (string, bool) {
	original, exists := urls[short]
	return original, exists
}

func UrlShortener(c *gin.Context) {
	var body struct {
		OriginURL string `json:"originURL"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty field",
		})
		return
	}

	result := shortener.GenerateShortUrl(body.OriginURL)
	urls[result] = body.OriginURL

	c.JSON(http.StatusOK, gin.H{
		"message": "Your short url - http://localhost:8080/" + result,
	})
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortcode")

	original, exists := GetOriginalURL(shortCode)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Shotrt URL not found" + original,
		})
		return
	}
	c.Redirect(http.StatusFound, original)

}
