package main

import (
	"url-shortener/internal/handlers"
	"url-shortener/internal/kafka"
	"url-shortener/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers.InitKafka()
	go kafka.StartConsumer()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to URL SHORTENER SERVICE, your entry point is http://localhost:8080/shortener/create",
		})
	})
	routers.UrlShortenerRouter(r)
	r.GET("/:shortcode", handlers.RedirectURL)
	r.Run()
}
