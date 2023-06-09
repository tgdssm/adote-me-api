package handlers

import (
	"api/helpers"
	"api/internal/core/domain"
	"api/internal/core/ports"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

type profileImageHandler struct {
	profileImageUseCase ports.ProfileImageUseCase
}

func NewProfileImageHandler(profileImageUseCase ports.ProfileImageUseCase, router *httprouter.Router) {
	handler := &profileImageHandler{
		profileImageUseCase: profileImageUseCase,
	}

	router.POST("/profile-image", Logger(Authenticator(handler.Create)))
	router.PUT("/profile-image/:id", Logger(Authenticator(handler.Update)))
}

func (ph profileImageHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(100); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}

	userID, err := strconv.ParseUint(r.FormValue("user_id"), 10, 64)
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

	file, _, err := r.FormFile("image")
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	filePath, fileName, _, err := helpers.GetFilePathAndFileName("profile-images")
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	profileImage := &domain.ProfileImage{
		FileName: fileName,
		FilePath: filePath,
		UserID:   userID,
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
	if err := r.ParseMultipartForm(100); err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}

	param, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userID, err := strconv.ParseUint(r.FormValue("user_id"), 10, 64)
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

	file, _, err := r.FormFile("image")
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	profileImage := &domain.ProfileImage{
		ID:       uint64(param),
		FileName: r.FormValue("file_name"),
		FilePath: r.FormValue("file_path"),
		UserID:   uint64(userID),
	}
	if helpers.CheckIfExists(strings.Replace(profileImage.FilePath, "\\\\", "\\", -1)) {
		helpers.DeleteFile(profileImage.FilePath)
	}

	filePath, fileName, _, err := helpers.GetFilePathAndFileName("profile-images")
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	profileImage.FileName = fileName
	profileImage.FilePath = filePath

	profileImage, err = ph.profileImageUseCase.Update(profileImage)
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
