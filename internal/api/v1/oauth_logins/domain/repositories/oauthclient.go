package repositories

import (
	"accounts/internal/infrastucture/oauth/google/entities"
	"accounts/internal/utils"
)

type OauthClientRepository interface {
	GetLink() utils.Result[string]
	GetToken(code string) utils.Result[string]
	GetUserInfo(accessToken string) utils.Result[entities.UserInfo]
}
