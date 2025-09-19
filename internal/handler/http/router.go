package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(router *chi.Mux, vehicleHandler *VehicleHandler) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/vehicles/add", vehicleHandler.Create)
	router.Put("/vehicles/{id}", vehicleHandler.Update)
}
