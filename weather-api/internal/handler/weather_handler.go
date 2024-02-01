package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

type WeatherHandler struct {
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (h *WeatherHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/weather/history", h.GetHistoricalWeather)
	app.Get("/api/weather/forecast", h.GetForecastWeather)
}

func (h *WeatherHandler) GetHistoricalWeather(ctx *fiber.Ctx) {
	fmt.Println(ctx)
	file, err := os.ReadFile("weather.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx.SendBytes(file)
}

func (h *WeatherHandler) GetForecastWeather(ctx *fiber.Ctx) {
	fmt.Println(ctx)
	file, err := os.ReadFile("weather.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx.SendBytes(file)
}

func (h *WeatherHandler) PredictWeather(ctx *fiber.Ctx) {

}
