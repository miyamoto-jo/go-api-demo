package user

import "github.com/miyamoto-jo/go-api-demo/datasource/user"

func NewDBRepository() Repository {
	return &dbRepository{}
}

type dbRepository struct {
	// Connection Info
}

func (m *dbRepository) GetUserInfo(userID string) (*user.UserInfo, error) {
	// ほんとはDBから取得する処理を書く
	return &user.UserInfo{
		ID:     userID,
		Name:   `miyamoto-jo`,
		Age:    30,
		Gender: `man`,
		Email:  `hogehoge@gmail.com`,
	}, nil
}
