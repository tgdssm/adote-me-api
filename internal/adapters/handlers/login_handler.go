package handlers

import (
	"api/helpers"
	"api/internal/core/ports"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type loginHandler struct {
	loginUseCase ports.LoginUseCase
}

func NewLoginHandler(loginUseCase ports.LoginUseCase, router *httprouter.Router) {
	handler := &loginHandler{
		loginUseCase: loginUseCase,
	}

	router.POST("/login", handler.Login)
}

func (lh loginHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	request := map[string]string{}

	if err = json.Unmarshal(body, &request); err != nil || request["email"] == "" || request["passwd"] == "" {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := lh.loginUseCase.GetByEmail(request["email"])
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = helpers.CheckPasswd(user.Passwd, request["passwd"]); err != nil {
		helpers.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := helpers.CreateToken(user.ID)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, token)

}
