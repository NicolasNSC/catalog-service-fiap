package repository

import (
	"context"

	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
)

type VehicleRepository interface {
	Save(ctx context.Context, vehicle *domain.Vehicle) error
}
