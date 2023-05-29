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

type petHandler struct {
	petUseCase      ports.PetUseCase
	petPhotoUseCase ports.PetPhotoUseCase
}

func NewPetHandler(petUseCase ports.PetUseCase, petPhotoUseCase ports.PetPhotoUseCase, router *httprouter.Router) {
	handler := &petHandler{
		petUseCase:      petUseCase,
		petPhotoUseCase: petPhotoUseCase,
	}

	router.POST("/pets", Logger(handler.Create))
	router.GET("/pets", Logger(Authenticator(handler.List)))
	router.GET("/pets/:id", Logger(Authenticator(handler.Get)))
}

func (ph petHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseMultipartForm(100)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}
	age, err := strconv.ParseUint(r.FormValue("age"), 10, 8)
	weight, err := strconv.ParseFloat(r.FormValue("weight"), 10)
	userID, err := strconv.ParseUint(r.FormValue("user_id"), 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	files := r.MultipartForm.File["photo"] // multiplos arquivos
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	pet := &domain.Pet{
		Name:         r.FormValue("pet_name"),
		Age:          age,
		Weight:       weight,
		Requirements: r.FormValue("requirements"),
		UserId:       userID,
	}

	pet, err = ph.petUseCase.Create(pet)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	for key, _ := range files {
		filePath, fileName, err := helpers.GetFilePathAndFileName("pet-images")
		if err != nil {
			helpers.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		petPhoto := &domain.PetPhoto{
			FileName: fileName,
			FilePath: filePath,
			PetID:    pet.ID,
		}

		petPhoto, err = ph.petPhotoUseCase.Create(petPhoto)
		if err != nil {
			helpers.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		pet.Photos = append(pet.Photos, *petPhoto)

		file, err := files[key].Open()
		if err != nil {
			helpers.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		defer file.Close()
		err = helpers.StorePictureInLocalFolder(file, "pet-images", filePath)

		if err != nil {
			helpers.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	helpers.JSON(w, http.StatusCreated, pet)
}

func (ph petHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := strings.ToLower(r.URL.Query().Get("pets"))
	var pets []domain.Pet

	pets, err := ph.petUseCase.List(query)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, pets)

}

func (ph petHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	petID, err := strconv.ParseUint(p.ByName("id"), 10, 64)

	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		helpers.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	pet, err := ph.petUseCase.Get(int(petID))
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, pet)
}
