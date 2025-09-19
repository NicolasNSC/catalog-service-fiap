package utils

import (
	"errors"
	"time"
)

func ValidateVehicleFields(brand, model string, year int, price float64) error {
	if brand == "" {
		return errors.New("brand cannot be empty")
	}
	if model == "" {
		return errors.New("model cannot be empty")
	}
	if year <= 1950 || year > time.Now().Year()+1 {
		return errors.New("vehicle year is invalid")
	}
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
