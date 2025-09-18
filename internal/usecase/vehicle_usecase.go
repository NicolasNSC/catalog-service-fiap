package usecase

import (
	"context"
	"time"

	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
	"github.com/NicolasNSC/catalog-service-fiap/internal/repository"
)

type InputCreateVehicleDTO struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

type OutputCreateVehicleDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type VehicleUseCaseInterface interface {
	Create(ctx context.Context, input InputCreateVehicleDTO) (*OutputCreateVehicleDTO, error)
}

type vehicleUseCase struct {
	repo repository.VehicleRepository
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUseCaseInterface {
	return &vehicleUseCase{
		repo: repo,
	}
}

func (v *vehicleUseCase) Create(ctx context.Context, input InputCreateVehicleDTO) (*OutputCreateVehicleDTO, error) {
	vehicle, err := domain.NewVehicle(input.Brand, input.Model, input.Color, input.Year, input.Price)
	if err != nil {
		return nil, err
	}

	err = v.repo.Save(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	output := &OutputCreateVehicleDTO{
		ID:        vehicle.ID,
		CreatedAt: vehicle.CreatedAt.Format(time.RFC3339),
	}

	return output, nil
}
