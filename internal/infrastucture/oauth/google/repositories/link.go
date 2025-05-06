package repositories

import "accounts/internal/utils"

func (r *OAuthGoogleRepository) GetLink() utils.Result[string] {
	return utils.Result[string]{
		Data: "https://accounts.google.com/o/oauth2/v2/auth?client_id=CLIENT_ID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE",
	}

}
