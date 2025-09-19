package repository

import (
	"context"

	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
)

//go:generate mockgen -source=vehicle_repository.go -destination=./mocks/vehicle_repository_mock.go -package=mocks
type VehicleRepository interface {
	Save(ctx context.Context, vehicle *domain.Vehicle) error
	GetByID(ctx context.Context, id string) (*domain.Vehicle, error)
	Update(ctx context.Context, vehicle *domain.Vehicle) error
}
