package controllers

import (
	"crypto/rand"
	"encoding/base64"
)

type GeneratorAPIKey struct {
	KeyLengthBytes int
}

func NewGeneratorAPIKey(keyLengthBytes int) *GeneratorAPIKey {
	return &GeneratorAPIKey{KeyLengthBytes: keyLengthBytes}
}

// GenerateAPIKey creates a new random API key of 384 bits, base64 URL encoded.
func (g *GeneratorAPIKey) GenerateAPIKey() (string, error) {
	b := make([]byte, g.KeyLengthBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
