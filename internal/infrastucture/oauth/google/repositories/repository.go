package repositories

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type OAuthGoogleRepository struct {
	restyClient     *resty.Client
	ClientIDiOS     string
	ClientIDAndroid string
	ClientIDWeb     string
	ClientSecret    string
	RedirectURI     string
}

func NewOAuthGoogleRepository(
	clientIDWeb, clientSecret, redirectURI string,
) *OAuthGoogleRepository {
	client := resty.New()
	client.SetBaseURL("https://oauth2.googleapis.com")
	client.SetTimeout(30 * time.Second) // Timeout para las peticiones

	return &OAuthGoogleRepository{
		restyClient:  client,
		ClientIDWeb:  clientIDWeb,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}
}
