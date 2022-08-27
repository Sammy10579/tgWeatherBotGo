package App

import (
	"fmt"
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

}
