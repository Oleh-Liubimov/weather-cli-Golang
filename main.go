package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		FeelslikeC float64 `json:"feelslike_c"`
	} `json:"current"`
}

func main() {

	fmt.Println("Hello, it is a cli weather app, provide a city where you want to know current weather:\n")
	var city string
	fmt.Scanln(&city)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	query := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no&alerts=no", apiKey, city)

	res, err := http.Get(query)

	if err != nil {
		fmt.Println("Error ocurred", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Weather API is not available", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error ocurred", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Failed to unmarshal json-file", err)
	}

	// fmt.Println(string(body))

	location, current := weather.Location, weather.Current

	fmt.Printf("%s,%s: Current temperature %.0f ℃, Condition: %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text)

	fmt.Scanln()

}
