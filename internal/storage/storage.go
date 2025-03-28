package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // No password set
	DB:       0,  // Use default DB
	Protocol: 2,  // Connection protocol
})

func SaveShortUrl(originalUrl, shortUrl string) (string, error) {
	UrlAlreadyExist, err := client.HGet(ctx, "original_to_short", originalUrl).Result()
	if err == nil {
		return UrlAlreadyExist, nil
	}

	err = client.HSet(ctx, "short_urls", shortUrl, originalUrl).Err()
	if err != nil {
		return "", fmt.Errorf("failed to save short URL: %v", err)
	}
	err = client.HSet(ctx, "original_to_short", originalUrl, shortUrl).Err()
	if err != nil {
		return "", fmt.Errorf("faied to save originalURL: %v", err)
	}

	return shortUrl, nil
}

func GetOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := client.HGet(ctx, "short_urls", shortUrl).Result()
	if err != nil {
		return "", fmt.Errorf("failed to obtain original URL: %v", err)
	}
	return originalUrl, nil
}
