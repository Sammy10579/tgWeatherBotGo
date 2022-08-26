package service

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) Get(city string) {
	w, err := owm.NewCurrent("F", "ru", os.Getenv("OPENWEATHER_APITOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	w.CurrentByName(city)
	fmt.Println(w)
}
