package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (r *postgresVehicleRepository) GetByID(ctx context.Context, id string) (*domain.Vehicle, error) {
	query := `SELECT id, brand, model, year, color, price, created_at, updated_at FROM vehicles WHERE id = $1`

	var v domain.Vehicle
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&v.ID, &v.Brand, &v.Model, &v.Year, &v.Color, &v.Price, &v.CreatedAt, &v.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("vehicle not found")
		}
		return nil, err
	}

	return &v, nil
}

func (r *postgresVehicleRepository) Update(ctx context.Context, vehicle *domain.Vehicle) error {
	query := `UPDATE vehicles 
	          SET brand = $1, model = $2, year = $3, color = $4, price = $5, updated_at = $6
	          WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query,
		vehicle.Brand,
		vehicle.Model,
		vehicle.Year,
		vehicle.Color,
		vehicle.Price,
		vehicle.UpdatedAt,
		vehicle.ID,
	)

	return err
}
