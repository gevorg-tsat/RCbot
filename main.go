package main

import (
	"github.com/joho/godotenv"
	"log"
	"untitledPetProject/internal/db"
)

func main() {
	err := godotenv.Load("secrets/.env")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}

}
