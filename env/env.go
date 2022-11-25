package env

import (
	"github.com/joho/godotenv"
 	"log"
	"os"
)

func GoEnvVariables(key string) string {
	//load .env file
	if err := godotenv.Load(".env"); err != nil{
		log.Fatal(".env loading error")
	}
	return os.Getenv(key)
}