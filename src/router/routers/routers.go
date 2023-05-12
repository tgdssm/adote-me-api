package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Uri          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request, httprouter.Params)
	AuthRequired bool
}

func Config(router *httprouter.Router) *httprouter.Router {
	routers := userRoutes

	for _, route := range routers {
		switch route.Method {
		case http.MethodPost:
			router.POST(route.Uri, route.Function)
		case http.MethodGet:
			router.GET(route.Uri, route.Function)
		case http.MethodPut:
			router.PUT(route.Uri, route.Function)
		case http.MethodDelete:
			router.DELETE(route.Uri, route.Function)
		}
	}

	return router
}
