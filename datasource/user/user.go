package user

type UserInfo struct {
	ID     string
	Name   string
	Age    int
	Gender string
	Email  string
}

// 18歳以上か
func (m *UserInfo) IsAdult() bool {
	if m.Age >= 18 {
		return true
	}
	return false
}
