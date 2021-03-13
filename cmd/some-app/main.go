package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")

	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error on load .env")
		}
	}

	initHttpService().Init()
}
