package controllers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2idCrypto implementa un hasher/verificador genérico con Argon2id + pepper.
// Formato de hash:
//
//	argon2id$v=19$m=<memKB>,t=<time>,p=<threads>$<salt_b64url>$<hash_b64url>
type Argon2idCrypto struct {
	Pepper       []byte
	Time         uint32
	MemoryKB     uint32
	Threads      uint8
	KeyLen       uint32
	SaltLenBytes int
}

func NewArgon2idCrypto(pepper []byte) *Argon2idCrypto {
	return &Argon2idCrypto{
		Pepper:       pepper,
		Time:         3,
		MemoryKB:     128 * 1024, // 128 MB
		Threads:      4,
		KeyLen:       32,
		SaltLenBytes: 16,
	}
}

func (c *Argon2idCrypto) Hash(content string) (string, error) {
	salt := make([]byte, c.SaltLenBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("salt generation: %w", err)
	}
	input := append(c.Pepper, []byte(content)...)
	sum := argon2.IDKey(input, salt, c.Time, c.MemoryKB, c.Threads, c.KeyLen)

	meta := fmt.Sprintf("argon2id$v=19$m=%d,t=%d,p=%d", c.MemoryKB, c.Time, c.Threads)
	saltB64 := base64.RawURLEncoding.EncodeToString(salt)
	sumB64 := base64.RawURLEncoding.EncodeToString(sum)
	return fmt.Sprintf("%s$%s$%s", meta, saltB64, sumB64), nil
}

func (c *Argon2idCrypto) Verify(storedHash, content string) bool {
	alg, memKB, t, threads, salt, expected, ok := parseArgon2Hash(storedHash)
	if !ok || alg != "argon2id" {
		return false
	}
	input := append(c.Pepper, []byte(content)...)
	sum := argon2.IDKey(input, salt, t, memKB, threads, uint32(len(expected)))
	return subtle.ConstantTimeCompare(sum, expected) == 1
}

// ── helpers ───────────────────────────────────────────

// Expresión regular para el formato estándar de Argon2id
var argon2Regex = regexp.MustCompile(`^(argon2id)\$v=\d+\$m=(\d+),t=(\d+),p=(\d+)\$([^$]+)\$([^$]+)$`)

// parseArgon2Hash extrae parámetros, salt y hash de un string Argon2id válido.

// parseArgon2Hash parsea el formato:
//
//	argon2id$v=19$m=<memKB>,t=<time>,p=<threads>$<salt_b64>$<hash_b64>
//
// Devuelve ok=false si el string no cumple el formato.
func parseArgon2Hash(s string) (alg string, memKB uint32, time uint32, threads uint8, salt, hash []byte, ok bool) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, "$")
	// Deben ser 5 partes: ["argon2id", "v=19", "m=...,t=...,p=...", "<salt>", "<hash>"]
	if len(parts) != 5 {
		return
	}
	if parts[0] != "argon2id" {
		return
	}
	alg = parts[0]

	// parts[1] debería ser "v=19" (lo toleramos pero no lo usamos)
	// parts[2] contiene "m=...,t=...,p=..."
	params := parts[2]
	if !strings.Contains(params, "m=") || !strings.Contains(params, "t=") || !strings.Contains(params, "p=") {
		return
	}

	// Extraer m,t,p sin depender del orden exacto
	for _, tok := range strings.Split(params, ",") {
		tok = strings.TrimSpace(tok)
		if kv := strings.SplitN(tok, "=", 2); len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
			switch key {
			case "m":
				if n, err := strconv.Atoi(val); err == nil {
					memKB = uint32(n)
				}
			case "t":
				if n, err := strconv.Atoi(val); err == nil {
					time = uint32(n)
				}
			case "p":
				if n, err := strconv.Atoi(val); err == nil && n >= 0 && n <= 255 {
					threads = uint8(n)
				}
			}
		}
	}
	if memKB == 0 || time == 0 || threads == 0 {
		return
	}

	// Decodificar salt y hash (URL-safe sin padding; fallback a std si falla)
	var err error
	salt, err = base64.RawURLEncoding.DecodeString(parts[3])
	if err != nil {
		salt, err = base64.RawStdEncoding.DecodeString(parts[3])
		if err != nil {
			return
		}
	}
	hash, err = base64.RawURLEncoding.DecodeString(parts[4])
	if err != nil {
		hash, err = base64.RawStdEncoding.DecodeString(parts[4])
		if err != nil {
			return
		}
	}

	ok = true
	return
}
