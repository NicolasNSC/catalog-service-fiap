package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NicolasNSC/catalog-service-fiap/internal/domain"
	"github.com/NicolasNSC/catalog-service-fiap/internal/repository"
	"github.com/stretchr/testify/suite"
)

type PostgresVehicleRepositoryTestSuite struct {
	suite.Suite
}

func Test_PostgresVehicleRepository(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PostgresVehicleRepositoryTestSuite))
}

func (suite *PostgresVehicleRepositoryTestSuite) Test_Save() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresVehicleRepository(db)

	vehicle := &domain.Vehicle{
		ID:        "123",
		Brand:     "Toyota",
		Model:     "Corolla",
		Year:      2022,
		Color:     "Blue",
		Price:     25000.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.T().Run("should save vehicle successfully", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO vehicles").
			WithArgs(
				vehicle.ID,
				vehicle.Brand,
				vehicle.Model,
				vehicle.Year,
				vehicle.Color,
				vehicle.Price,
				vehicle.CreatedAt,
				vehicle.UpdatedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Save(context.Background(), vehicle)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	suite.T().Run("should return error when insert fails", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO vehicles").
			WithArgs(
				vehicle.ID,
				vehicle.Brand,
				vehicle.Model,
				vehicle.Year,
				vehicle.Color,
				vehicle.Price,
				vehicle.CreatedAt,
				vehicle.UpdatedAt,
			).
			WillReturnError(errors.New("insert error"))

		err = repo.Save(context.Background(), vehicle)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func (suite *PostgresVehicleRepositoryTestSuite) Test_GetByID() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresVehicleRepository(db)

	vehicle := &domain.Vehicle{
		ID:        "123",
		Brand:     "Toyota",
		Model:     "Corolla",
		Year:      2022,
		Color:     "Blue",
		Price:     25000.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.T().Run("should get vehicle by id successfully", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "brand", "model", "year", "color", "price", "created_at", "updated_at",
		}).AddRow(
			vehicle.ID,
			vehicle.Brand,
			vehicle.Model,
			vehicle.Year,
			vehicle.Color,
			vehicle.Price,
			vehicle.CreatedAt,
			vehicle.UpdatedAt,
		)

		mock.ExpectQuery("SELECT id, brand, model, year, color, price, created_at, updated_at FROM vehicles WHERE id = \\$1").
			WithArgs(vehicle.ID).
			WillReturnRows(rows)

		got, err := repo.GetByID(context.Background(), vehicle.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if got == nil || got.ID != vehicle.ID {
			t.Errorf("expected vehicle with ID %s, got %+v", vehicle.ID, got)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	suite.T().Run("should return error when vehicle not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, brand, model, year, color, price, created_at, updated_at FROM vehicles WHERE id = \\$1").
			WithArgs("notfound").
			WillReturnError(sql.ErrNoRows)

		got, err := repo.GetByID(context.Background(), "notfound")
		if err == nil || got != nil {
			t.Errorf("expected error and nil vehicle, got err=%v, got=%+v", err, got)
		}
	})

	suite.T().Run("should return error on query failure", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, brand, model, year, color, price, created_at, updated_at FROM vehicles WHERE id = \\$1").
			WithArgs("fail").
			WillReturnError(errors.New("query error"))

		got, err := repo.GetByID(context.Background(), "fail")
		if err == nil || got != nil {
			t.Errorf("expected error and nil vehicle, got err=%v, got=%+v", err, got)
		}
	})
}

func (suite *PostgresVehicleRepositoryTestSuite) Test_Update() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresVehicleRepository(db)

	vehicle := &domain.Vehicle{
		ID:        "123",
		Brand:     "Toyota",
		Model:     "Corolla",
		Year:      2022,
		Color:     "Blue",
		Price:     25000.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.T().Run("should update vehicle successfully", func(t *testing.T) {
		mock.ExpectExec("UPDATE vehicles").
			WithArgs(
				vehicle.Brand,
				vehicle.Model,
				vehicle.Year,
				vehicle.Color,
				vehicle.Price,
				vehicle.UpdatedAt,
				vehicle.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Update(context.Background(), vehicle)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	suite.T().Run("should return error when update fails", func(t *testing.T) {
		mock.ExpectExec("UPDATE vehicles").
			WithArgs(
				vehicle.Brand,
				vehicle.Model,
				vehicle.Year,
				vehicle.Color,
				vehicle.Price,
				vehicle.UpdatedAt,
				vehicle.ID,
			).
			WillReturnError(errors.New("update error"))

		err := repo.Update(context.Background(), vehicle)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
