package controller

import "net/http"

func HealthCheckController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`go api demo active`))
}
