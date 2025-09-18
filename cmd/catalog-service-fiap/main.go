package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/NicolasNSC/catalog-service-fiap/internal/handler/http"
	"github.com/NicolasNSC/catalog-service-fiap/internal/repository"
	"github.com/NicolasNSC/catalog-service-fiap/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	loadConfig()

	db := setupDatabase()
	defer db.Close()

	repo := repository.NewPostgresVehicleRepository(db)
	useCase := usecase.NewVehicleUseCase(repo)
	vehicleHandler := handler.NewVehicleHandler(useCase)

	router := setupRouter(vehicleHandler)

	startServer(router)
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
}

func setupDatabase() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Fatal: could not connect to database: %v", err)
	}
	if err = db.PingContext(context.Background()); err != nil {
		log.Fatalf("Fatal: could not ping database: %v", err)
	}
	return db
}

func setupRouter(vehicleHandler *handler.VehicleHandler) *chi.Mux {
	r := chi.NewRouter()
	handler.SetupRoutes(r, vehicleHandler)
	return r
}

func startServer(router *chi.Mux) {
	apiPort := os.Getenv("API_PORT")
	log.Printf("Info: server starting on port %s", apiPort)
	if err := http.ListenAndServe(":"+apiPort, router); err != nil {
		log.Fatalf("Fatal: could not start server: %v", err)
	}
}
