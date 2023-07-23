package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
