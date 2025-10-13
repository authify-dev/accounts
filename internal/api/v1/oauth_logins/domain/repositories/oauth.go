package repositories

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

type OAuthLoginRepository interface {
	Save(role entities.OAuthLogin) utils.Result[entities.OAuthLogin]
	Search(uuid string) (entities.OAuthLogin, error)
	SearchAll() ([]entities.OAuthLogin, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.OAuthLogin, error)
	View(data []entities.OAuthLogin)
}
