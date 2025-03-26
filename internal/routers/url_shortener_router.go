package routers

import (
	"url-shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UrlShortenerRouter(c *gin.Engine) {
	url_shortener := c.Group("/shortener")
	{
		url_shortener.POST("/create", handlers.UrlShortener)
	}
}
