package main

import (
	"RCbot/internal/bot"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("secrets/.env")
	if err != nil {
		log.Fatal(err)
	}
	if err = bot.Run(); err != nil {
		log.Fatal(err)
	}
}
