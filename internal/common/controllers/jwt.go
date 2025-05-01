package controllers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTController gestiona generación y validación de JWTs usando RSA (RS256).
// PrivateKey y PublicKey deben contener las claves en formato PEM.
type JWTController struct {
	PrivateKey string
	PublicKey  string
}

// parseRSAPrivateKey decodifica un PEM PKCS#1 o PKCS#8 en *rsa.PrivateKey.
func parseRSAPrivateKey(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("no se pudo decodificar PEM de clave privada")
	}
	// PKCS#1
	if priv, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return priv, nil
	}
	// PKCS#8
	keyIfc, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("falló parseo PKCS#8: %w", err)
	}
	priv, ok := keyIfc.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("clave PKCS#8 no es RSA")
	}
	return priv, nil
}

// parseRSAPublicKey decodifica un PEM en *rsa.PublicKey.
func parseRSAPublicKey(pemStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("no se pudo decodificar PEM de clave pública")
	}
	pubIfc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("falló parseo clave pública: %w", err)
	}
	pub, ok := pubIfc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("clave pública no es RSA")
	}
	return pub, nil
}

// GenerateToken crea un JWT con claims personalizados y lo firma con RS256.
func (j *JWTController) GenerateToken(data map[string]interface{}, expireSeconds int) (string, error) {
	privKey, err := parseRSAPrivateKey(j.PrivateKey)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	for k, v := range data {
		claims[k] = v
	}
	if expireSeconds > 0 {
		claims["exp"] = time.Now().Add(time.Duration(expireSeconds) * time.Second).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privKey)
}

// ValidateToken valida un JWT firmado con RS256 y retorna sus claims.
// Si el token ha expirado, devuelve un error con mensaje personalizado.
func (j *JWTController) ValidateToken(tokenString string) (map[string]interface{}, error) {
	pubKey, err := parseRSAPublicKey(j.PublicKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, fmt.Errorf("algoritmo inesperado: %s", token.Method.Alg())
		}
		return pubKey, nil
	})

	// Si hay un error al parsear, comprobamos si se debe a expiración
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, fmt.Errorf("%v", ve.Inner)
		}
		return nil, err
	}

	// Si el token es válido, retornamos las claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}
