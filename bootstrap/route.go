package bootstrap

import (
	"goblogCalmk/pkg/route"
	"goblogCalmk/routes"

	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
