package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString string
	Port             int64
	SecretKey        []byte
)

func LoadConfig() {
	var err error
	if err = godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.ParseInt(os.Getenv("API_PORT"), 10, 0)

	if err != nil {
		Port = 5000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"))
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
