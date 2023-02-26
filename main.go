package main

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/website/middleware"
	"github.com/ibilalkayy/website/routes"
)

func Execute() error {
	routes.Routes()
	fmt.Println("Starting the server at :4000")
	return http.ListenAndServe(":4000", nil)
}

func main() {
	err := Execute()
	middleware.HandleError(err)
}
