package handler

import (
	"github.com/gofiber/fiber"
	"github.com/shostyi-ve/apka/weather/internal/service/weather"
	"time"
)

type WeatherHandler struct {
	service *weather.WeatherService
}

type HistoricalWeatherRequest struct {
	City string `json:"city"`
}

func NewWeatherHandler(service *weather.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/weather/history", h.GetHistoricalWeather)
	app.Get("/api/weather/forecast", h.GetForecastWeather)
}

func (h *WeatherHandler) GetHistoricalWeather(ctx *fiber.Ctx) {
	timeStartStr, _ := time.Parse(time.RFC3339, ctx.Query("timeStart"))
	timeEndStr, _ := time.Parse(time.RFC3339, ctx.Query("timeEnd"))
	city := ctx.Query("city")

	history, err := h.service.GetHistory(ctx.Context(), timeStartStr, timeEndStr, city)
	if err != nil {
		ctx.Write(err)
	}

	ctx.SendBytes([]byte(history))
}

func (h *WeatherHandler) GetForecastWeather(ctx *fiber.Ctx) {
	city := ctx.Query("city")

	history, err := h.service.GetPrediction(ctx.Context(), city)
	if err != nil {
		ctx.Write(err)
	}

	ctx.SendBytes([]byte(history))
}
