package repositories

import (
	"accounts/internal/utils"
	"encoding/json"
	"fmt"
	"net/url"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (r *OAuthGoogleRepository) GetToken(code string) utils.Result[string] {

	decodedCode, err := url.QueryUnescape(code)
	if err != nil {
		return utils.Result[string]{Err: fmt.Errorf("failed to decode code: %w", err)}
	}

	resp, err := r.restyClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"client_id":     r.ClientIDWeb,
			"client_secret": r.ClientSecret,
			"code":          decodedCode,
			"redirect_uri":  r.RedirectURI,
			"grant_type":    "authorization_code",
		}).
		SetResult(&TokenResponse{}).
		Post("/token")

	if err != nil {
		return utils.Result[string]{
			Err: fmt.Errorf("failed to make request: %w", err),
		}
	}

	if resp.StatusCode() != 200 {
		// Parsear el cuerpo de la respuesta como GoogleErrorResponse
		var googleError map[string]any
		if parseErr := json.Unmarshal(resp.Body(), &googleError); parseErr != nil {
			return utils.Result[string]{
				Err: fmt.Errorf("failed to parse error response: %w", parseErr),
			}
		}

		// Retornar el error con la descripción específica
		return utils.Result[string]{
			Err: fmt.Errorf("error from Google API: %s", googleError["error"]),
		}
	}

	token := resp.Result().(*TokenResponse)

	// Devolvemos el access_token
	return utils.Result[string]{Data: token.AccessToken}
}
