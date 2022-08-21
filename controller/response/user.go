package response

type GetUserResponse struct {
	Name   string `json:"name"`
	Age    int    `json:""`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}
