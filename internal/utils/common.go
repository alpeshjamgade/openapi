package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
