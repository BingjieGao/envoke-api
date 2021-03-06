package main

import (
	"net/http"

	"github.com/Envoke-org/envoke-api/api"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// Create http router
	router := httprouter.New()

	// Create api and add routes
	api.NewApi().AddRoutes(router)

	// Start HTTP server with router
	http.ListenAndServe(":8888", router)
}
