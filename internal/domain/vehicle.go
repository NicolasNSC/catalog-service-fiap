package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID        string    `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Color     string    `json:"color"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVehicle(brand, model, color string, year int, price float64) (*Vehicle, error) {
	if brand == "" {
		return nil, errors.New("brand cannot be empty")
	}

	if model == "" {
		return nil, errors.New("model cannot be empty")
	}

	if year <= 1950 || year > time.Now().Year()+1 {
		return nil, errors.New("vehicle year is invalid")
	}

	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	return &Vehicle{
		ID:        uuid.New().String(),
		Brand:     brand,
		Model:     model,
		Year:      year,
		Color:     color,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
