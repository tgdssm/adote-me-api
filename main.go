package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	router := router.Generator()

	log.Fatal(http.ListenAndServe(":8080", router))
}
