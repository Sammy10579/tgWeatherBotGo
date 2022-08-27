package openweather

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type OpenWeather struct {
	o *OpenWeather
}

func NewOpenWeather() *OpenWeather {
	return &OpenWeather{}
}

func (o *OpenWeather) OWStart() error {
	w, err := owm.NewCurrent("F", "ru", os.Getenv("OPENWEATHER_APITOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	city := owm.City{}
	w.CurrentByName(city)
	fmt.Println(w)
}
