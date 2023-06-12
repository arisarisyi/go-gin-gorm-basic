package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func init(){
	LoadEnvVariables()
}

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error loading .env file")
	}
}