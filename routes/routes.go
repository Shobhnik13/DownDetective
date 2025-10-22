package routes

import (
	"down_detective/controllers"

	"github.com/gorilla/mux"
)

func RegisterAppRoutes(router *mux.Router) {
	router.HandleFunc("/ping-urls", controllers.PingURLsHandler).Methods("POST")	
}
