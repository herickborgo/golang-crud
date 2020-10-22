package router

import (
	"rest-api/controllers/probe"

	"github.com/gorilla/mux"
)

// GetRouter return routes
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/probe", probe.Index).Methods("GET")
	return router
}
