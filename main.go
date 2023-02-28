package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibilalkayy/website/middleware"
	"github.com/ibilalkayy/website/routes"
)

func Execute() error {
	routes.Routes()
	fmt.Println("Starting the server at :8080")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	err := Execute()
	middleware.HandleError(err)
}
