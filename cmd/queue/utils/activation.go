package utils

import (
	"accounts/internal/core/settings"
	"accounts/internal/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GenerateBodyActivation(username, code string) utils.Either[string] {
	url := settings.Settings.EMAIL_TEMPLATE_ACTIVATION_URL

	// Realizamos la solicitud GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error realizando el GET: %v\n", err)
		return utils.Either[string]{Err: err}
	}
	defer resp.Body.Close()

	// Leemos el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error leyendo la respuesta: %v\n", err)
		return utils.Either[string]{Err: err}
	}

	content := string(body)

	// Reemplazamos las variables del template
	content = strings.ReplaceAll(content, "{user_name}", username)
	content = strings.ReplaceAll(content, "{activation_code}", code)

	// Imprimimos el contenido obtenido
	return utils.Either[string]{Data: content}
}
