package handlers

import (
	"api/helpers"
	"api/internal/core/domain"
	"api/internal/core/ports"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(userUseCase ports.UserUseCase, router *httprouter.Router) {
	handler := &userHandler{
		userUseCase: userUseCase,
	}

	router.POST("/users", Logger(handler.Create))
	router.GET("/users", Logger(Authenticator(handler.List)))
	router.GET("/users/:id", Logger(Authenticator(handler.Get)))
	router.PUT("/users/:id", Logger(Authenticator(handler.Update)))
	router.DELETE("/users/:id", Logger(Authenticator(handler.Delete)))

}

func (uh userHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user *domain.User

	if err = json.Unmarshal(body, &user); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = user.Prepare(false); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err = uh.userUseCase.Create(user)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	user.Token, err = helpers.CreateToken(user.ID)

	helpers.JSON(w, http.StatusCreated, user)
}

func (uh userHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := strings.ToLower(r.URL.Query().Get("users"))
	var users []domain.User

	users, err := uh.userUseCase.List(query)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, users)

}

func (uh userHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID, err := strconv.ParseUint(p.ByName("id"), 10, 64)

	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		helpers.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	user, err := uh.userUseCase.Get(int(userID))
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, user)
}

func (uh userHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	userID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if tokenUserID != userID {
		helpers.ERROR(w, http.StatusForbidden, errors.New("it is not possible to update a user other than the one who is logged in"))
		return
	}

	var user *domain.User

	if err = json.Unmarshal(body, &user); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.ID = userID

	if err = user.Prepare(true); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = uh.userUseCase.Update(user)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, user)
}

func (uh userHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if tokenUserID != userID {
		helpers.ERROR(w, http.StatusForbidden, errors.New("it is not possible to update a user other than the one who is logged in"))
		return
	}

	if err = uh.userUseCase.Delete(int(userID)); err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
	}

	helpers.JSON(w, http.StatusNoContent, nil)
}
