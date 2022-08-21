package request

type GetUserRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
