package repositories

import (
	"accounts/internal/infrastucture/oauth/google/entities"
	"accounts/internal/utils"
	"errors"
)

func (g *OAuthGoogleRepository) GetUserInfo(accessToken string) utils.Result[entities.UserInfo] {
	resp, err := g.restyClient.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetResult(&entities.UserInfo{}).
		Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return utils.Result[entities.UserInfo]{Err: err}
	}

	if resp.StatusCode() != 200 {
		return utils.Result[entities.UserInfo]{
			Err: errors.New("failed to get user info"),
		}
	}

	userInfo := resp.Result().(*entities.UserInfo)

	return utils.Result[entities.UserInfo]{Data: *userInfo}
}
