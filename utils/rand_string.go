package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[source.Intn(len(charset))]
	}

	return string(result)
}
