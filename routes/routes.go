package routes

import (
	"net/http"

	"github.com/ibilalkayy/website/controllers"
	"github.com/ibilalkayy/website/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.ErrorHandling(controllers.Home))
}
