package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

type WeatherHandler struct{}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (h *WeatherHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/weather", h.GetWeather)
}

func (h *WeatherHandler) GetWeather(ctx *fiber.Ctx) {
	fmt.Println(ctx)
	file, err := os.ReadFile("weather.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx.SendBytes(file)
}
