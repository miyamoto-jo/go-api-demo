package controller

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/miyamoto-jo/go-api-demo/controller/parser"
	"github.com/miyamoto-jo/go-api-demo/controller/resultcode"
)

func CreateEndPoints() *mux.Router {
	r := mux.NewRouter().SkipClean(true)
	for url, resource := range *LoadServices() {
		r.HandleFunc(url, resource.f).Methods(resource.method)
	}
	return r
}

const (
	URIHealthCheck = `/healthcheck`

	URIGetUserInfo = `/user`
)

// ハンドラとメソッドをまとめた構造体
type resource struct {
	f      func(w http.ResponseWriter, r *http.Request)
	method string
}

func LoadServices() *map[string]*resource {
	services := make(map[string]*resource)

	services[URIHealthCheck] = &resource{f: HealthCheckController, method: http.MethodGet}

	services[URIGetUserInfo] = &resource{f: panicInterceptor(GetUser), method: http.MethodPost}

	return &services
}

func panicInterceptor(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stackTrace := make([]byte, 2048)
				len := runtime.Stack(stackTrace, true)
				msg := fmt.Sprintf(`[Err]panic %v %s`, err, string(stackTrace[:len]))
				parser.JsonResponseError(w, resultcode.CodeErrorPanicFatal, errors.New(msg))
			}
		}()
		f(w, r)
	}
}
