package handlers

import (
	"context"
	"log"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafka() {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9093"),
		Topic:    "shortened-urls",
		Balancer: &kafka.LeastBytes{},
	}
}

func UrlShortener(c *gin.Context) {
	var body struct {
		OriginalURL string `json:"originalURL"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty field",
		})
		return
	}

	short_url := shortener.GenerateShortUrl(body.OriginalURL)
	_, err := storage.SaveShortUrl(body.OriginalURL, short_url)
	if err != nil {
		log.Printf("failed to save URL: %v", err)
	}

	err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(body.OriginalURL),
		Value: []byte(short_url),
	})
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your short url - http://localhost:8080/" + short_url,
	})
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortcode")

	original, err := storage.GetOriginalUrl(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Shotrt URL not found",
		})
		return
	}

	c.Redirect(http.StatusFound, original)
}
