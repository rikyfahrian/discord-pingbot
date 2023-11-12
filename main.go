package main

import (
	"log"
	"ping-bot/bot"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("failed to load env")
	}

}

func main() {

	bot.Start()

}
