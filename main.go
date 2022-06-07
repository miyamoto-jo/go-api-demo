package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/miyamoto-jo/go-api-demo/controller"
)

func main() {
	if err := launchServer(controller.CreateEndPoints()); err != nil {
		fmt.Println("fatal error on launch server: ", err)
	}
	defer terminate()
}

func terminate() {
	fmt.Println("terminate処理")
}

func launchServer(r http.Handler) error {
	server := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", 8080),
		WriteTimeout: 80 * time.Second,
		ReadTimeout:  80 * time.Second,
	}
	return server.ListenAndServe()
}
