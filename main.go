package main

import (
	"github.com/joho/godotenv"
	"log"
	"untitledPetProject/internal/bot"
)

func main() {
	err := godotenv.Load("secrets/.env")
	if err != nil {
		log.Fatal(err)
	}
	err = bot.Run()
	if err != nil {
		log.Fatal(err)
	}
}
