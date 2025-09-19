package usecase

import (
	"context"
	"time"

	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
	"github.com/NicolasNSC/catalog-service-fiap/internal/repository"
	"github.com/NicolasNSC/catalog-service-fiap/internal/utils"
	"github.com/google/uuid"
)

//go:generate mockgen -source=vehicle_usecase.go -destination=./mocks/vehicle_usecase_mock.go -package=mocks
type VehicleUseCaseInterface interface {
	Create(ctx context.Context, input dto.InputCreateVehicleDTO) (*dto.OutputCreateVehicleDTO, error)
	Update(ctx context.Context, id string, input dto.InputUpdateVehicleDTO) error
}

type vehicleUseCase struct {
	repo repository.VehicleRepository
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUseCaseInterface {
	return &vehicleUseCase{
		repo: repo,
	}
}

func (v *vehicleUseCase) Create(ctx context.Context, input dto.InputCreateVehicleDTO) (*dto.OutputCreateVehicleDTO, error) {
	err := utils.ValidateVehicleFields(input.Brand, input.Model, input.Year, input.Price)
	if err != nil {
		return nil, err
	}

	vehicle := &domain.Vehicle{
		ID:        uuid.New().String(),
		Brand:     input.Brand,
		Model:     input.Model,
		Year:      input.Year,
		Color:     input.Color,
		Price:     input.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = v.repo.Save(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	output := &dto.OutputCreateVehicleDTO{
		ID:        vehicle.ID,
		CreatedAt: vehicle.CreatedAt.Format(time.RFC3339),
	}

	return output, nil
}

func (uc *vehicleUseCase) Update(ctx context.Context, id string, input dto.InputUpdateVehicleDTO) error {
	vehicle, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = utils.ValidateVehicleFields(input.Brand, input.Model, input.Year, input.Price)
	if err != nil {
		return err
	}

	vehicle.Brand = input.Brand
	vehicle.Model = input.Model
	vehicle.Color = input.Color
	vehicle.Year = input.Year
	vehicle.Price = input.Price
	vehicle.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, vehicle)
}
