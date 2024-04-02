package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file Error Message %s", err.Error())
	}

}
