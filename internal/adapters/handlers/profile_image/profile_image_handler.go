package profile_image

import (
	"api/helpers"
	"api/internal/core/ports"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type profileImageHandler struct {
	profileImageUseCase ports.ProfileImageUseCase
}

func NewProfileImageHandler(profileImageUseCase ports.ProfileImageUseCase, router *httprouter.Router) {
	handler := &profileImageHandler{
		profileImageUseCase: profileImageUseCase,
	}

	router.POST("/image-profile", handler.Create)
	router.POST("/image-profile/:id", handler.Update)
}

func (ph profileImageHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(100); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}

}

func (ph profileImageHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
