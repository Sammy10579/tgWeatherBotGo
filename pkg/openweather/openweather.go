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

func (o *OpenWeather) ByCity(city string) (answer string, err error) {
	err = o.o.CurrentByName(city)
	if err != nil {
		fmt.Println("Введите пожалуйста существующий город")
		return
	}
	answer = "Температура в г. " + city + " " + strconv.Itoa(int(o.o.Main.Temp))
	return
}

func (o *OpenWeather) ByLocation(long, lat float64) (answer string, err error) {

	err = o.o.CurrentByCoordinates(&owm.Coordinates{
		Longitude: long,
		Latitude:  lat,
	})
	if err != nil {
		fmt.Println("Нет такой локации")
		return "", err
	}
	answer = "Температура в " + o.o.Name + " " + strconv.Itoa(int(o.o.Main.Temp))
	fmt.Println(long, lat)
	return
}

func (o *OpenWeather) Start() error {
	log.Println("OpenWeather api start")
	return nil
}
