package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"tgWeatherBotGo/pkg/openweather"
	"tgWeatherBotGo/pkg/telegam"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	openWeather := openweather.NewOpenWeather()

	telegramBot := telegam.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
