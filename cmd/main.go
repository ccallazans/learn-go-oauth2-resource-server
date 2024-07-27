package main

import (
	"log"

	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := api.NewRouter()
	router.Logger.Fatal(router.Start(":1323"))
}
