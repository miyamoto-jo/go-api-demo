package controller

import (
	"net/http"

	"github.com/gorilla/mux"
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
)

// ハンドラとメソッドをまとめた構造体
type resource struct {
	f      func(w http.ResponseWriter, r *http.Request)
	method string
}

func LoadServices() *map[string]*resource {
	services := make(map[string]*resource)

	services[URIHealthCheck] = &resource{f: HealthCheckController, method: http.MethodGet}

	return &services
}
