package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationSettings struct{}

func NewApplicationSettings() *ApplicationSettings {
	err := godotenv.Load(".env")
	if err != nil {
		// Just log an error - the environment file is not required
		log.Println("No valid .env file found")
	}
	return &ApplicationSettings{}
}

func (s *ApplicationSettings) IsDev() bool {
	return os.Getenv("ENVIRONMENT_TYPE") == "DEV"
}

func (s *ApplicationSettings) ServerPort() string {
	return os.Getenv("PORT")
}
