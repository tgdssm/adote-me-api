package main

import (
	"api/helpers"
	handlers "api/internal/adapters/handlers/user"
	repositories "api/internal/adapters/repositories/user"
	"api/internal/core/ports"
	"api/internal/core/usecases"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	helpers.LoadConfig()
	fmt.Println(helpers.ConnectionString)
}

func main() {
	router := httprouter.New()

	var userRepo ports.UserRepository = repositories.NewUserMysqlRepository()
	userUseCase := usecases.NewUserUseCase(userRepo)
	handlers.NewUserHandler(userUseCase, router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", helpers.Port), router))
}
