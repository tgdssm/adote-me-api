package router

import (
	"api/src/router/routes"

	"github.com/julienschmidt/httprouter"
)

func Generator() *httprouter.Router {
	router := httprouter.New()
	return routes.Config(router)
}
