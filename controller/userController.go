package controller

import (
	"net/http"

	"github.com/miyamoto-jo/go-api-demo/controller/parser"
	"github.com/miyamoto-jo/go-api-demo/controller/request"
	"github.com/miyamoto-jo/go-api-demo/controller/response"
	"github.com/miyamoto-jo/go-api-demo/controller/resultcode"
	"github.com/miyamoto-jo/go-api-demo/service/user"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	req := request.GetUserRequest{}
	if err := parser.ParseAndValidateRequest(r, &req); err != nil {
		parser.JsonResponseError(w, resultcode.CodeRequestParseFailure, err)
		return
	}

	service := user.New(req.UserID)
	ui, err := service.GetUserInfo()
	if err != nil {
		parser.JsonResponseError(w, resultcode.CodeRequestParseFailure, err)
		return
	}

	res := response.GetUserResponse{
		Name:   ui.Name,
		Age:    ui.Age,
		Gender: ui.Gender,
		Email:  ui.Email,
	}

	parser.JsonResponse(w, res)
}
