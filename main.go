package main

import (
	"crisecheverria/weather/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Missing WEATHERAPI_KEY")
		os.Exit(1)
	}

	defaultLocation := "Gothenburg"
	location := defaultLocation

	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?appid=%s&units=metric&q=%s", apiKey, location)

	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching weather data: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Weather API returned non-200 status: %d\n", res.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading weather data: %v\n", err)
		os.Exit(1)
	}

	var weatherData types.WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		panic(err)
	}

	groupedWeather := make(map[string][]string)

	currentDate := time.Now().Format("2006-01-02") // Format the date as "YYYY-MM-DD"

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for _, entry := range weatherData.List {
		dateTimeParts := strings.Split(entry.DtTxt, " ")
		date := dateTimeParts[0]
		time := dateTimeParts[1]
		temp := entry.Main.Temp

		var tempString string

		if temp <= 10 {
			tempString = cyan(fmt.Sprintf("%.0f°C", temp))
		} else if temp >= 11 && temp <= 20 {
			tempString = yellow(fmt.Sprintf("%.0f°C", temp))
		} else {
			tempString = red(fmt.Sprintf("%.0f°C", temp))
		}
		if date == currentDate {
			formattedEntry := fmt.Sprintf("%s, Temp: %s, Min: %.0f°C, Max: %.0f°C, Feels Like: %.0f°C, Weather: %s",
				time[:5],
				tempString,
				entry.Main.TempMin,
				entry.Main.TempMax,
				entry.Main.FeelsLike,
				entry.Weather[0].Description,
			)

			groupedWeather[date] = append(groupedWeather[date], formattedEntry)
		}
	}

	for date, entries := range groupedWeather {
		fmt.Printf(location+": %s\n", date)
		for _, entry := range entries {
			fmt.Println(entry)
		}
		fmt.Println()
	}
}
