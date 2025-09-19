package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/NicolasNSC/catalog-service-fiap/docs"
)

func SetupRoutes(router *chi.Mux, vehicleHandler *VehicleHandler) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Post("/vehicles/add", vehicleHandler.Create)
	router.Put("/vehicles/{id}", vehicleHandler.Update)
}
