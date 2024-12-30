package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	fmt.Println("Hello, World!")
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=574f9deba16b4951bb712025242912&q=Cali&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status code is not 200: %d", res.StatusCode))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour


	fmt.Printf("%s, %s: %.1f°C, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)
	fmt.Println("Hourly forecast:")
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		fmt.Printf("Time: %s, Temperature: %.1f°C, Chance of rain: %.1f%%\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain)
	}
}
