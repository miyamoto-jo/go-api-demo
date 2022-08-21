package user

import "github.com/miyamoto-jo/go-api-demo/datasource/user"

type Repository interface {
	GetUserInfo(userID string) (*user.UserInfo, error)
}
