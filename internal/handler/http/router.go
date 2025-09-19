package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(router *chi.Mux, vehicleHandler *VehicleHandler) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/vehicles", func(r chi.Router) {
		r.Post("/add", vehicleHandler.Create)
		r.Put("/{id}", vehicleHandler.Update)
	})
}
