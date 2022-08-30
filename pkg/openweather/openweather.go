package openweather

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type OpenWeather struct {
	o   *owm.CurrentWeatherData
	bot *tgbotapi.BotAPI
}

func NewOpenWeather(data *owm.CurrentWeatherData) *OpenWeather {
	return &OpenWeather{o: data}
}

func (o *OpenWeather) ByCity(city string) (answer string, err error) {
	err = o.o.CurrentByName(city)
	if err != nil {
		fmt.Println(err)
		return "Введите пожалуйста существующий город", err
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
		fmt.Println(err)
		return "Нет такой локации", err
	}
	answer = "Температура в " + o.o.Name + " " + strconv.Itoa(int(o.o.Main.Temp))
	fmt.Println(long, lat)
	return
}

func (o *OpenWeather) Weather(message *tgbotapi.Message) (answer string, err error) {
	if message != nil {
		log.Printf("[%s] %s", message.From.UserName, message.Text)

		if message.Location != nil {
			answer, err = o.ByLocation(message.Location.Longitude, message.Location.Latitude)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			answer, err = o.ByCity(message.Text)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

func (o *OpenWeather) Start() error {
	log.Println("OpenWeather api start")
	return nil
}
