package repository

import (
	"context"
	"database/sql"

	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
)

type postgresVehicleRepository struct {
	db *sql.DB
}

func NewPostgresVehicleRepository(db *sql.DB) VehicleRepository {
	return &postgresVehicleRepository{
		db: db,
	}
}

func (r *postgresVehicleRepository) Save(ctx context.Context, vehicle *domain.Vehicle) error {
	query := `INSERT INTO vehicles (id, brand, model, year, color, price, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		vehicle.ID,
		vehicle.Brand,
		vehicle.Model,
		vehicle.Year,
		vehicle.Color,
		vehicle.Price,
		vehicle.CreatedAt,
		vehicle.UpdatedAt,
	)

	return err
}
