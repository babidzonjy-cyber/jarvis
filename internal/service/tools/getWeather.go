package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GeocodingResponse struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

type WeatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
	} `json:"current"`
}

func GetWeather(city, lang string) (string, error) {
	cityEncoded := url.QueryEscape(city)
	httpURLforCoordinates := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=%s&format=json", cityEncoded, lang)

	respOfCoordinates, err := http.Get(httpURLforCoordinates)
	if err != nil {
		return "", fmt.Errorf("http get: %w", err)
	}
	defer respOfCoordinates.Body.Close()

	var geo GeocodingResponse
	if err := json.NewDecoder(respOfCoordinates.Body).Decode(&geo); err != nil {
		return "", fmt.Errorf("geocoding json decode: %w", err)
	}

	lat := geo.Results[0].Latitude
	long := geo.Results[0].Longitude

	httpURLforTemp := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,wind_speed_10m", lat, long)

	respOfTemp, err := http.Get(httpURLforTemp)
	if err != nil {
		return "", fmt.Errorf("weather http get: %w", err)
	}
	defer respOfTemp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(respOfTemp.Body).Decode(&weather); err != nil {
		return "", fmt.Errorf("json decode: %w", err)
	}

	return fmt.Sprintf(
		"Температура: %.1f°C, ветер: %.1f км/ч",
		weather.Current.Temperature,
		weather.Current.WindSpeed,
	), nil
}
