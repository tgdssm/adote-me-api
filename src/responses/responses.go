package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}
