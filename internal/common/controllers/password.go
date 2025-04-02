package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordController se encarga de manejar el cifrado y verificación de contraseñas utilizando un secreto personal (pepper).
type PasswordController struct {
	secret string
}

// NewPasswordController crea una nueva instancia de PasswordController con el secreto personal.
func NewPasswordController(secret string) PasswordController {
	return PasswordController{secret: secret}
}

// HashPassword recibe una contraseña en texto plano, le añade el secreto personal y retorna su versión cifrada.
func (pc PasswordController) HashPassword(password string) (string, error) {
	// Combinar la contraseña con el secreto (pepper)
	passwordWithPepper := password + pc.secret
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword compara una contraseña en texto plano (a la que se le añade el secreto) con una contraseña cifrada.
// Retorna true si la contraseña en texto plano (más el secreto) corresponde a la versión cifrada.
func (pc PasswordController) CheckPassword(password, hashedPassword string) bool {
	passwordWithPepper := password + pc.secret
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithPepper))
	return err == nil
}
