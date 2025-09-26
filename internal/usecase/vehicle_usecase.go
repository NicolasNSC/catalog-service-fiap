package usecase

import (
	"context"
	"log"
	"time"

	"github.com/NicolasNSC/catalog-service-fiap/internal/client"
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
	repo           repository.VehicleRepository
	showcaseClient client.ShowcaseClientInterface
}

func NewVehicleUseCase(repo repository.VehicleRepository, showcaseClient client.ShowcaseClientInterface) VehicleUseCaseInterface {
	return &vehicleUseCase{
		repo:           repo,
		showcaseClient: showcaseClient,
	}
}

func (vuc *vehicleUseCase) Create(ctx context.Context, input dto.InputCreateVehicleDTO) (*dto.OutputCreateVehicleDTO, error) {
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

	err = vuc.repo.Save(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	listingDTO := dto.CreateListingDTO{
		VehicleID: vehicle.ID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Price:     vehicle.Price,
	}

	err = vuc.showcaseClient.CreateListing(ctx, listingDTO)
	if err != nil {
		log.Printf("Warning: failed to notify showcase-service about new vehicle %s: %v", vehicle.ID, err)
	}

	output := &dto.OutputCreateVehicleDTO{
		ID:        vehicle.ID,
		CreatedAt: vehicle.CreatedAt.Format(time.RFC3339),
	}

	return output, nil
}

func (vuc *vehicleUseCase) Update(ctx context.Context, id string, input dto.InputUpdateVehicleDTO) error {
	vehicle, err := vuc.repo.GetByID(ctx, id)
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

	err = vuc.repo.Update(ctx, vehicle)
	if err != nil {
		return err
	}

	listingDTO := dto.UpdateListingDTO{
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Price: vehicle.Price,
	}

	err = vuc.showcaseClient.UpdateListing(ctx, vehicle.ID, listingDTO)
	if err != nil {
		log.Printf("Warning: failed to notify showcase-service about vehicle update %s: %v", vehicle.ID, err)
	}

	return nil
}
