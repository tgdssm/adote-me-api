package main

import (
	"api/helpers"
	"api/internal/adapters/handlers"
	"api/internal/adapters/repositories"
	"api/internal/core/ports"
	"api/internal/core/usecases"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func init() {
	// Gerando secret key
	//key := make([]byte, 64)
	//if _, err := rand.Read(key); err != nil {
	//	log.Fatal(key)
	//}
	//
	//base64String := base64.StdEncoding.EncodeToString(key)

	helpers.LoadConfig()
	fmt.Println(helpers.ConnectionString)
}

func main() {
	router := httprouter.New()

	var userRepo ports.UserRepository = repositories.NewUserMysqlRepository()
	var userUseCase ports.UserUseCase = usecases.NewUserUseCase(userRepo)
	handlers.NewUserHandler(userUseCase, router)

	var profileImageRepo ports.ProfileImageRepository = repositories.NewProfileImageMysqlRepository()
	var profileImageUseCase ports.ProfileImageUseCase = usecases.NewProfileImageUseCase(profileImageRepo)
	handlers.NewProfileImageHandler(profileImageUseCase, router)

	var loginRepo ports.LoginRepository = repositories.NewLoginMysqlRepository()
	var loginUseCase ports.LoginUseCase = usecases.NewLoginUseCase(loginRepo)
	handlers.NewLoginHandler(loginUseCase, userUseCase, router)

	var petRepo ports.PetRepository = repositories.NewPetMysqlRepository()
	var petUseCase ports.PetUseCase = usecases.NewPetUseCase(petRepo)
	var petPhotoRepo ports.PetPhotoRepository = repositories.NewPetPhotoMysqlRepository()
	var petPhotoUseCase ports.PetPhotoUseCase = usecases.NewPetPhotoUseCase(petPhotoRepo)
	handlers.NewPetHandler(petUseCase, petPhotoUseCase, router)

	handlers.NewFileServerHandler(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", helpers.Port), router))
}
