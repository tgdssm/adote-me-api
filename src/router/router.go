package router

import (
	"api/src/router/routers"

	"github.com/julienschmidt/httprouter"
)

func Generator() *httprouter.Router {
	router := httprouter.New()
	return routers.Config(router)
}
