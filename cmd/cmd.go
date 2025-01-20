package cmd

import (
	"fmt"
	"go-be-commerce/delivery/http"
)

func Execute() {
	// start init container

	// start queue service

	// start http service
	http := http.ServeHttp()
	http.Listen(fmt.Sprintf(":%d", 3000))
}
