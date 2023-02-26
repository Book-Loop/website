package routes

import (
	"net/http"

	"github.com/ibilalkayy/website/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Home)
}
