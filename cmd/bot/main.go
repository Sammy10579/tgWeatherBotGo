package main

import (
	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	App2 "tgWeatherBotGo/App"
	"tgWeatherBotGo/pkg/openweather"
	"tgWeatherBotGo/pkg/telegam"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	w, err := owm.NewCurrent("C", "ru", os.Getenv("OPENWEATHER_APITOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	openWeather := openweather.NewOpenWeather(w)
	if err := openWeather.Start(); err != nil {
		log.Fatalln(err)
	}

	telegramBot := telegam.NewBot(bot)
	if err := telegramBot.Start(openWeather.ByCity, openWeather.ByLocation); err != nil {
		log.Fatal(err)
	}

	App := App2.NewApplication(openWeather, telegramBot)
	App.Run()
}
