package repositories

import (
	"accounts/internal/utils"
	"fmt"
	"net/url"
)

func (r *OAuthGoogleRepository) GetLink() utils.Result[string] {

	baseURL := "https://accounts.google.com/o/oauth2/v2/auth"

	params := url.Values{}
	params.Add("client_id", r.ClientIDWeb)
	params.Add("redirect_uri", r.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "email profile")
	params.Add("access_type", "offline")
	params.Add("prompt", "consent")

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	return utils.Result[string]{
		Data: url,
	}

}
