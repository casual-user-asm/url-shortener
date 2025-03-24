package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func GenerateShortUrl(originURL string) string {
	hash := sha256.Sum256([]byte(originURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	return strings.TrimRight(encoded[:8], "=")
}
