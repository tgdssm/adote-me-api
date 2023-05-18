package handlers

import (
	"api/helpers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Logger(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r, p)
	}
}

func Authenticator(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := helpers.ValidateToken(r); err != nil {
			helpers.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r, p)
	}
}
