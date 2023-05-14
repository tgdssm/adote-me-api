package handlers

import (
	"api/helpers"
	"api/internal/core/domain"
	"api/internal/core/ports"
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
	router.POST("/users/:id", handler.Get)

}

func (uh userHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(100); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := &domain.User{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		Passwd:    r.FormValue("passwd"),
		Cellphone: r.FormValue("cellphone"),
	}

	file, hfile, err := r.FormFile("profilePicture")
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	user.PicturePath, err = helpers.StorePictureInLocalFolder(file, hfile, "profile-pictures")
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
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
