package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Criando usuário!"))
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Buscando usuários!"))
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Buscando usuário!"))
}

func PutUser(w http.ResponseWriter, r *http.Request, _p httprouter.Params) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Deletando usuário!"))
}
