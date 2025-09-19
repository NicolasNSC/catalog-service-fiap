package utils_test

import (
	"testing"
	"time"

	"github.com/NicolasNSC/catalog-service-fiap/internal/utils"
)

func TestValidateVehicleFields(t *testing.T) {
	tests := []struct {
		name    string
		brand   string
		model   string
		year    int
		price   float64
		wantErr bool
	}{
		{
			name:    "valid fields",
			brand:   "Toyota",
			model:   "Corolla",
			year:    time.Now().Year(),
			price:   20000,
			wantErr: false,
		},
		{
			name:    "empty brand",
			brand:   "",
			model:   "Civic",
			year:    2020,
			price:   15000,
			wantErr: true,
		},
		{
			name:    "empty model",
			brand:   "Honda",
			model:   "",
			year:    2020,
			price:   15000,
			wantErr: true,
		},
		{
			name:    "year too old",
			brand:   "Ford",
			model:   "Mustang",
			year:    1940,
			price:   30000,
			wantErr: true,
		},
		{
			name:    "year in the future",
			brand:   "Tesla",
			model:   "Model S",
			year:    time.Now().Year() + 2,
			price:   80000,
			wantErr: true,
		},
		{
			name:    "zero price",
			brand:   "Chevrolet",
			model:   "Onix",
			year:    2022,
			price:   0,
			wantErr: true,
		},
		{
			name:    "negative price",
			brand:   "BMW",
			model:   "320i",
			year:    2022,
			price:   -10000,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateVehicleFields(tt.brand, tt.model, tt.year, tt.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateVehicleFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
