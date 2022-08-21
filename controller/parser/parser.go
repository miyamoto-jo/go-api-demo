package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	vv "github.com/go-playground/validator/v10"
)

func ParseAndValidateRequest(r *http.Request, bodyStruct interface{}) error {
	var err error
	if err = parseRequestBody(r, bodyStruct); err != nil {
		return err
	}

	if err = validateRequestStructure(bodyStruct); err != nil {
		return err
	}
	return nil
}

// リクエストのPOSTボディを指定構造体に格納する
func parseRequestBody(r *http.Request, bodyStruct interface{}) error {
	if reflect.ValueOf(bodyStruct).Type().Kind() != reflect.Ptr {
		panic(`parse request argment must be pointer.`)
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	// read body data parse json
	err = json.Unmarshal(b, bodyStruct)
	if err != nil {
		return err
	}
	return nil
}

// 構造体をvalidateタグの記述に基づいてチェックする
func validateRequestStructure(bodyStruct interface{}) error {
	validator := vv.New()
	return validator.Struct(bodyStruct)
}
