package parser

import (
	"encoding/json"
	"net/http"

	"github.com/miyamoto-jo/go-api-demo/controller/resultcode"
)

func JsonResponse(w http.ResponseWriter, obj interface{}) {
	err, ok := obj.(error)
	if ok {
		JsonResponseError(w, resultcode.CodeErrorFailure, err)
		return
	}
	JsonResponseFull(w, resultcode.CodeSuccess, obj, nil)
}

func JsonResponseError(w http.ResponseWriter, rc resultcode.ResultCode, err error) {
	if rc == resultcode.CodeSuccess {
		panic(`cann't set CodeSuccess`)
	}
	JsonResponseFull(w, rc, nil, err)
}

type basicResponse struct {
	ResultCode int         `json:"result_code"`
	Data       interface{} `json:"data"`
	ErrorMsg   *string     `json:"error"`
}

func JsonResponseFull(w http.ResponseWriter, rc resultcode.ResultCode, obj interface{}, err error) {
	var errmsg *string
	if err != nil {
		msg := err.Error()
		errmsg = &msg
	}
	base := &basicResponse{ResultCode: int(rc), Data: obj, ErrorMsg: errmsg}
	bin, err := json.Marshal(base)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set(`Content-Type`, `application/json`)
	w.Header().Set(`Cache-Control`, `no-store`) // 脆弱性対応
	w.Write(bin)
}
