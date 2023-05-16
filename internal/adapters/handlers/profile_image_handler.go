package handlers

import (
	"api/helpers"
	"api/internal/core/domain"
	"api/internal/core/ports"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type profileImageHandler struct {
	profileImageUseCase ports.ProfileImageUseCase
}

func NewProfileImageHandler(profileImageUseCase ports.ProfileImageUseCase, router *httprouter.Router) {
	handler := &profileImageHandler{
		profileImageUseCase: profileImageUseCase,
	}

	router.POST("/profile-image", handler.Create)
	router.POST("/profile-image/:id", handler.Update)
}

func (ph profileImageHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(100); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}

	userID, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	filePath, fileName, err := helpers.GetFilePathAndFileName("profile-images")
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	profileImage := &domain.ProfileImage{
		FileName: fileName,
		FilePath: filePath,
		UserID:   uint64(userID),
	}

	profileImage, err = ph.profileImageUseCase.Create(profileImage)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = helpers.StorePictureInLocalFolder(file, "profile-images", filePath)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, profileImage)

}

func (ph profileImageHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
