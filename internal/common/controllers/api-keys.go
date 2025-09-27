package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const defaultLen = 48 // 384 bits

type GeneratorAPIKey struct {
	KeyLengthBytes int
	Environment    string // "live" | "test"
}

func NewGeneratorAPIKey(n int, env string) *GeneratorAPIKey {
	if n <= 0 {
		n = defaultLen
	}
	if env == "" {
		env = "live"
	}
	return &GeneratorAPIKey{KeyLengthBytes: n, Environment: env}
}

// KeyMaterial: salida del generador (sin dependencias de dominio)
type KeyMaterial struct {
	KeyID          string // no-secreto (HEX), útil para lookup directo
	Prefix         string // prefijo del tramo aleatorio de la SECRET (para índices)
	SecretKey      string // "sk_<env>_<keyid>_<randomBase64Url>"
	PublishableKey string // "pk_<env>_<keyid>_<randomBase64Url>"
	Environment    string
}

// GeneratePair crea ambas llaves embeddiendo env y keyid en el formato acordado.
func (g *GeneratorAPIKey) GeneratePair() (KeyMaterial, error) {
	keyID, err := genHexID(12) // ~96 bits, en HEX (24 chars)
	if err != nil {
		return KeyMaterial{}, err
	}

	secretRand, err := genBase64URL(g.KeyLengthBytes)
	if err != nil {
		return KeyMaterial{}, err
	}
	pubRand, err := genBase64URL(g.KeyLengthBytes)
	if err != nil {
		return KeyMaterial{}, err
	}

	secret := fmt.Sprintf("sk_%s_%s_%s", g.Environment, keyID, secretRand)
	public := fmt.Sprintf("pk_%s_%s_%s", g.Environment, keyID, pubRand)

	return KeyMaterial{
		KeyID:          keyID,
		Prefix:         firstN(secretRand, 8),
		SecretKey:      secret,
		PublishableKey: public,
		Environment:    g.Environment,
	}, nil
}

// ── Helpers ────────────────────────────────────────────────────

func genBase64URL(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("rand.Read: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil // URL-safe, sin padding
}

func genHexID(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("rand.Read: %w", err)
	}
	return hex.EncodeToString(b), nil // solo [0-9a-f] → no contiene "_"
}

func firstN(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if len(s) < n {
		return s
	}
	return s[:n]
}

func IsSecretKey(s string) bool      { return strings.HasPrefix(s, "sk_") }
func IsPublishableKey(s string) bool { return strings.HasPrefix(s, "pk_") }
