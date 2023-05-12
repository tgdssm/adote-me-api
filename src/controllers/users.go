package controllers

import (
	"api/src/db"
	"api/src/helpers"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(100); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Name:   r.FormValue("name"),
		Email:  r.FormValue("email"),
		Passwd: r.FormValue("passwd"),
	}

	file, hfile, err := r.FormFile("profilePicture")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	user.PicturePath, err = helpers.StorePictureInLocalFolder(file, hfile, "profile-pictures")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.StartConnection()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := strings.ToLower(r.URL.Query().Get("users"))
	db, err := db.StartConnection()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.GetUsers(query)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param, err := strconv.ParseInt(p.ByName("id"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.StartConnection()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, err := repository.GetUser(int(param))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}

func PutUser(w http.ResponseWriter, r *http.Request, _p httprouter.Params) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Deletando usuário!"))
}
