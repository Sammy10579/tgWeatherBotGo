package openweather

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"strconv"
)

type OpenWeather struct {
	o *owm.CurrentWeatherData
}

func NewOpenWeather(data *owm.CurrentWeatherData) *OpenWeather {
	return &OpenWeather{o: data}
}

func (o *OpenWeather) PrintWeather(city string) (answer string, err error) {
	err = o.o.CurrentByName(city)
	if err != nil {
		fmt.Println("Введите пожалуйста существующий город")
		return
	}
	answer = "Температура в г. " + city + " " + strconv.Itoa(int(o.o.Main.Temp))
	return
}

func (o *OpenWeather) Start() error {
	log.Println("OpenWeather api start")
	return nil
}
