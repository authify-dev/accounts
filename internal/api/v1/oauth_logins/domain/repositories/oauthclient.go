package repositories

import "accounts/internal/utils"

type OauthClientRepository interface {
	GetLink() utils.Result[string]
}
