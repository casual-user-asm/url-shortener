package main

import (
	"url-shortener/internal/handlers"
	"url-shortener/internal/kafka"
	"url-shortener/internal/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers.InitKafka()
	go kafka.StartConsumer()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to URL SHORTENER SERVICE",
		})
	})

	routers.UrlShortenerRouter(r)
	r.GET("/:shortcode", handlers.RedirectURL)
	r.Run(":8080")
}
