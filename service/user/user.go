package user

import (
	duser "github.com/miyamoto-jo/go-api-demo/datasource/user"
	ruser "github.com/miyamoto-jo/go-api-demo/repository/user"
)

type Service interface {
	GetUserInfo() (*duser.UserInfo, error)
}

func New(userID string) Service {
	return &service{
		repo:   ruser.NewDBRepository(),
		userID: userID,
	}
}

type service struct {
	repo ruser.Repository

	userID string
}

func (m *service) GetUserInfo() (*duser.UserInfo, error) {
	return m.repo.GetUserInfo(m.userID)
}
