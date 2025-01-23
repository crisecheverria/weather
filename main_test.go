package main

import (
	"crisecheverria/weather/types"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestProcessWeatherData(t *testing.T) {

	// Sample JSON response (you can add more test cases with different data)
	sampleJSON := `
	{
		"list": [
			{"dt_txt": "2024-04-26 12:00:00", "main": {"temp": 20, "feels_like": 18, "temp_min": 15, "temp_max": 22}, "weather": [{"description": "Sunny"}]},
			{"dt_txt": "2024-04-26 15:00:00", "main": {"temp": 22, "feels_like": 20, "temp_min": 18, "temp_max": 25}, "weather": [{"description": "Cloudy"}]},
			{"dt_txt": "2024-04-27 09:00:00", "main": {"temp": 18, "feels_like": 16, "temp_min": 12, "temp_max": 20}, "weather": [{"description": "Rainy"}]}
		]
	}`

	var weatherData types.WeatherResponse
	err := json.Unmarshal([]byte(sampleJSON), &weatherData)
	if err != nil {
		t.Fatalf("Error unmarshalling sample JSON: %v", err)
	}

	groupedWeather := make(map[string][]string)

	for _, entry := range weatherData.List {
		dateTimeParts := strings.Split(entry.DtTxt, " ")
		date := dateTimeParts[0]
		time := dateTimeParts[1]
		formattedEntry := fmt.Sprintf("%s, Temp: %.0f°C, Min: %.0f°C, Max: %.0f°C, Feels Like: %.0f°C, Weather: %s",
			time[:5],
			entry.Main.Temp,
			entry.Main.TempMin,
			entry.Main.TempMax,
			entry.Main.FeelsLike,
			entry.Weather[0].Description,
		)
		groupedWeather[date] = append(groupedWeather[date], formattedEntry)
	}

	// Assertions (add more assertions as needed)
	if len(groupedWeather) != 2 {
		t.Errorf("Expected 2 dates, got %d", len(groupedWeather))
	}

	if len(groupedWeather["2024-04-26"]) != 2 {
		t.Errorf("Expected 2 entries for 2024-04-26, got %d", len(groupedWeather["2024-04-26"]))

	}
}
