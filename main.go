package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Loader()
	fmt.Println(config.ConnectionString)
}

func main() {
	router := router.Generator()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
