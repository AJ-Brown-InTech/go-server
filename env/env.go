package env

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GoEnvVariables(key string) string {
	//load .env file
	
	if err := godotenv.Load(".env"); err != nil{
		log.Fatal(".env loading error")
	}
	return os.Getenv(key)
}