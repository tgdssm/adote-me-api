package handlers

import (
	"api/helpers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func NewFileServerHandler(router *httprouter.Router) {
	router.GET("/pet-images/:image", FileServer)
}

func FileServer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	imageName := p.ByName("image")
	pathDir, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}
	absolutePath := filepath.Join(fmt.Sprintf("%s/%s", pathDir, "pet-images"), imageName)
	file, err := os.Open(absolutePath)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fileModTime, err := os.Stat(absolutePath)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	http.ServeContent(w, r, "", fileModTime.ModTime(), file)
}
