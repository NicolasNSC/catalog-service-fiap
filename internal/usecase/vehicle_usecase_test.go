package usecase_test

import (
	"context"
	"testing"

	mclient "github.com/NicolasNSC/catalog-service-fiap/internal/client/mocks"
	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
	"github.com/NicolasNSC/catalog-service-fiap/internal/repository/mocks"
	"github.com/NicolasNSC/catalog-service-fiap/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type VehicleUseCaseSuite struct {
	suite.Suite

	ctx            context.Context
	repository     *mocks.MockVehicleRepository
	showcaseClient *mclient.MockShowcaseClientInterface
}

func (suite *VehicleUseCaseSuite) BeforeTest(_, _ string) {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	suite.ctx = context.Background()
	suite.repository = mocks.NewMockVehicleRepository(ctrl)
	suite.showcaseClient = mclient.NewMockShowcaseClientInterface(ctrl)
}

func Test_VehicleUseCaseSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(VehicleUseCaseSuite))
}

func (suite *VehicleUseCaseSuite) Test_Create() {
	input := dto.InputCreateVehicleDTO{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2022,
		Color: "White",
		Price: 100000,
	}

	suite.T().Run("should create a vehicle successfully", func(t *testing.T) {
		suite.repository.EXPECT().Save(suite.ctx, gomock.Any()).Return(nil)
		suite.showcaseClient.EXPECT().CreateListing(suite.ctx, gomock.Any()).Return(nil)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		output, err := usecase.Create(suite.ctx, input)
		suite.NoError(err)
		suite.NotNil(output)
		suite.NotEmpty(output.ID)
		suite.NotEmpty(output.CreatedAt)
	})

	suite.T().Run("should return error when input validation fails", func(t *testing.T) {
		input := dto.InputCreateVehicleDTO{
			Brand: "",
			Model: "",
			Year:  0,
			Color: "",
			Price: 0,
		}

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		output, err := usecase.Create(suite.ctx, input)
		suite.Error(err)
		suite.Nil(output)
	})

	suite.T().Run("should return error when repository save fails", func(t *testing.T) {
		suite.repository.EXPECT().
			Save(gomock.Any(), gomock.Any()).
			Return(assert.AnError)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		output, err := usecase.Create(suite.ctx, input)
		suite.Error(err)
		suite.Nil(output)
	})

	suite.T().Run("should log warning when showcase client fails but still create vehicle", func(t *testing.T) {
		suite.repository.EXPECT().Save(suite.ctx, gomock.Any()).Return(nil)
		suite.showcaseClient.EXPECT().CreateListing(suite.ctx, gomock.Any()).Return(assert.AnError)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		output, err := usecase.Create(suite.ctx, input)
		suite.NoError(err)
		suite.NotNil(output)
		suite.NotEmpty(output.ID)
		suite.NotEmpty(output.CreatedAt)
	})

}

func (suite *VehicleUseCaseSuite) Test_Update() {
	id := "vehicle-123"
	input := dto.InputUpdateVehicleDTO{
		Brand: "Ford",
		Model: "Focus",
		Year:  2023,
		Color: "Blue",
		Price: 120000,
	}
	existingVehicle := &domain.Vehicle{
		ID:    id,
		Brand: "Ford",
		Model: "Fiesta",
		Year:  2020,
		Color: "Red",
		Price: 80000,
	}

	suite.T().Run("should update a vehicle successfully", func(t *testing.T) {
		suite.repository.EXPECT().
			GetByID(suite.ctx, id).
			Return(existingVehicle, nil)
		suite.repository.EXPECT().
			Update(suite.ctx, gomock.Any()).
			Return(nil)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		err := usecase.Update(suite.ctx, id, input)
		suite.NoError(err)
	})

	suite.T().Run("should return error when input validation fails", func(t *testing.T) {
		input := dto.InputUpdateVehicleDTO{
			Brand: "",
			Model: "",
			Year:  0,
			Color: "",
			Price: 0,
		}

		suite.repository.EXPECT().
			GetByID(suite.ctx, id).
			Return(existingVehicle, nil)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		err := usecase.Update(suite.ctx, id, input)
		suite.Error(err)
	})

	suite.T().Run("should return error when vehicle not found", func(t *testing.T) {
		suite.repository.EXPECT().
			GetByID(suite.ctx, id).
			Return(nil, assert.AnError)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		err := usecase.Update(suite.ctx, id, input)
		suite.Error(err)
	})

	suite.T().Run("should return error when repository update fails", func(t *testing.T) {
		suite.repository.EXPECT().
			GetByID(suite.ctx, id).
			Return(existingVehicle, nil)
		suite.repository.EXPECT().
			Update(suite.ctx, gomock.Any()).
			Return(assert.AnError)

		usecase := usecase.NewVehicleUseCase(suite.repository, suite.showcaseClient)
		err := usecase.Update(suite.ctx, id, input)
		suite.Error(err)
	})
}
