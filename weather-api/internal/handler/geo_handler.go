package handler

import (
	"github.com/shostyi-ve/apka/weather/pkg/countriesnow"

	"github.com/gofiber/fiber"
)

type GeoHandler struct {
	countryNowClient *countriesnow.Client
}

func NewGeoHandler(countryNowClient *countriesnow.Client) *GeoHandler {
	return &GeoHandler{
		countryNowClient: countryNowClient,
	}
}

func (h *GeoHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/countries", h.GetCountries)
	app.Get("/api/cities/:country", h.GetCities)
}

func (h *GeoHandler) GetCountries(ctx *fiber.Ctx) {
	countries, err := h.countryNowClient.AllCountries(ctx.Context())
	if err != nil {
		panic(err)
	}

	ctx.JSON(countries)
}

func (h *GeoHandler) GetCities(ctx *fiber.Ctx) {
	req := countriesnow.CitiesByCountryRequest{
		CountryName: ctx.Params("country"),
	}

	cities, err := h.countryNowClient.CitiesByCountry(ctx.Context(), req)
	if err != nil {
		panic(err)
	}

	ctx.JSON(cities)
}

type countryDto struct {
	Name string `json:"name"`
}
