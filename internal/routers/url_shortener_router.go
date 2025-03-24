package routers

import (
	"url-shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func url_shortener_router(c *gin.Engine) {
	url_shortener := c.Group("/shortener")
	{
		url_shortener.POST("/create", handlers.UrlShortener)
	}
}
