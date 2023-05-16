package handlers

import (
	"api/helpers"
	"api/internal/core/domain"
	"api/internal/core/ports"
	"encoding/json"
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

	router.POST("/users", handler.Create)
	router.GET("/users", handler.List)
	router.GET("/users/:id", handler.Get)
	router.PUT("/users/:id", handler.Update)

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

	if err = user.Prepare(); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err = uh.userUseCase.Create(user)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, user)
}

func (uh userHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := strings.ToLower(r.URL.Query().Get("users"))

	users, err := uh.userUseCase.List(query)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, users)

}

func (uh userHandler) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	param, err := strconv.ParseInt(p.ByName("id"), 10, 32)

	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := uh.userUseCase.Get(int(param))
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

	param, err := strconv.ParseInt(p.ByName("id"), 10, 32)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user *domain.User

	if err = json.Unmarshal(body, &user); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.ID = uint64(param)

	if err = user.Prepare(); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err = uh.userUseCase.Update(user)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, user)
}
