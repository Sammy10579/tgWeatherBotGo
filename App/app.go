package App

import (
	"fmt"
	"log"
	"tgWeatherBotGo/pkg/openweather"
	"tgWeatherBotGo/pkg/telegam"
)

type Application struct {
	Weather *openweather.OpenWeather
	Bot     *telegam.Bot
}

func NewApplication(Weather *openweather.OpenWeather, Bot *telegam.Bot) *Application {
	return &Application{
		Weather: Weather,
		Bot:     Bot,
	}
}

func (a *Application) Run() {
	fmt.Println("Run Application")
	a.Bot.MassageHandler(a.Weather.Weather)

	if err := a.Bot.Start(); err != nil {
		log.Fatal(err)
	}
}
