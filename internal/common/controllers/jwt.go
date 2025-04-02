package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTController struct {
	PublicKey  string // No se utiliza en llaves simétricas
	PrivateKey string // Se utiliza como llave secreta
}

func (j *JWTController) GenerateToken(data map[string]interface{}, expired int) (string, error) {
	// Creamos un MapClaims para incluir los datos y la expiración
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}

	// Agregamos la expiración, si se especifica
	if expired > 0 {
		claims["exp"] = time.Now().Add(time.Duration(expired) * time.Second).Unix()
	}

	// Creamos el token usando el método de firma HS256 (HMAC con SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmamos el token con la llave secreta (PrivateKey)
	tokenString, err := token.SignedString([]byte(j.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTController) ValidateToken(tokenString string) (map[string]interface{}, error) {
	// Parseamos el token, indicando cómo obtener la llave de verificación
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificamos que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("método de firma inesperado", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(j.PrivateKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Validamos que las claims sean del tipo MapClaims y que el token sea válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.NewValidationError("token inválido", jwt.ValidationErrorSignatureInvalid)
}
