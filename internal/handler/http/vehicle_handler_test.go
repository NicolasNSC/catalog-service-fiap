package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
	h "github.com/NicolasNSC/catalog-service-fiap/internal/handler/http"
	"github.com/NicolasNSC/catalog-service-fiap/internal/usecase/mocks"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type VehicleHandlerSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *mocks.MockVehicleUseCaseInterface
	handler *h.VehicleHandler
}

func (suite *VehicleHandlerSuite) BeforeTest(_, _ string) {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	suite.ctx = context.Background()
	suite.useCase = mocks.NewMockVehicleUseCaseInterface(ctrl)
	suite.handler = h.NewVehicleHandler(suite.useCase)

}

func Test_VehicleHandlerSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(VehicleHandlerSuite))
}

func (suite *VehicleHandlerSuite) Test_Create() {
	suite.T().Run("Create - Success", func(t *testing.T) {
		input := dto.InputCreateVehicleDTO{
			Brand: "Toyota",
			Model: "Corolla",
			Year:  2022,
			Color: "Blue",
			Price: 20000.00,
		}
		expectedOutput := &dto.OutputCreateVehicleDTO{
			ID:        "123",
			CreatedAt: "2023-10-01T12:00:00Z",
		}

		suite.useCase.EXPECT().Create(suite.ctx, input).Return(expectedOutput, nil)

		body, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/vehicles/add", bytes.NewReader(body))
		w := httptest.NewRecorder()

		suite.handler.Create(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusCreated, resp.StatusCode)
		var got dto.OutputCreateVehicleDTO
		err := json.NewDecoder(resp.Body).Decode(&got)
		suite.NoError(err)
		suite.Equal(*expectedOutput, got)
	})

	suite.T().Run("Create - Invalid Body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/vehicles/add", strings.NewReader("invalid-json"))
		w := httptest.NewRecorder()

		suite.handler.Create(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusBadRequest, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		suite.Contains(string(body), "Invalid request body")
	})

	suite.T().Run("Create - Use Case Error", func(t *testing.T) {
		input := dto.InputCreateVehicleDTO{
			Brand: "Toyota",
			Model: "Corolla",
			Year:  2022,
		}

		suite.useCase.EXPECT().
			Create(gomock.Any(), input).
			Return(&dto.OutputCreateVehicleDTO{}, errors.New("some error"))

		body, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/vehicles/add", bytes.NewReader(body))
		w := httptest.NewRecorder()

		suite.handler.Create(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		suite.Contains(string(bodyResp), "Failed to create vehicle")
	})
}

func (suite *VehicleHandlerSuite) Test_Update() {
	suite.T().Run("Update - Success", func(t *testing.T) {
		id := "123"
		input := dto.InputUpdateVehicleDTO{
			Brand: "Honda",
			Model: "Civic",
			Year:  2023,
			Color: "Red",
			Price: 25000.00,
		}

		suite.useCase.EXPECT().
			Update(gomock.Any(), id, input).
			Return(nil)

		body, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+id, bytes.NewReader(body))
		req = req.WithContext(suite.ctx)
		req = muxSetURLParam(req, "id", id)
		w := httptest.NewRecorder()

		suite.handler.Update(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusOK, resp.StatusCode)
	})

	suite.T().Run("Update - Missing ID", func(t *testing.T) {
		input := dto.InputUpdateVehicleDTO{
			Brand: "Honda",
			Model: "Civic",
			Year:  2023,
			Color: "Red",
			Price: 25000.00,
		}
		body, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/", bytes.NewReader(body))
		w := httptest.NewRecorder()

		suite.handler.Update(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusBadRequest, resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		suite.Contains(string(bodyResp), "VehicleID is required")
	})

	suite.T().Run("Update - Invalid Body", func(t *testing.T) {
		id := "123"
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+id, strings.NewReader("invalid-json"))
		req = muxSetURLParam(req, "id", id)
		w := httptest.NewRecorder()

		suite.handler.Update(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusBadRequest, resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		suite.Contains(string(bodyResp), "Invalid request body")
	})

	suite.T().Run("Update - Use Case Error", func(t *testing.T) {
		id := "123"
		input := dto.InputUpdateVehicleDTO{
			Brand: "Honda",
			Model: "Civic",
			Year:  2023,
			Color: "Red",
			Price: 25000.00,
		}

		suite.useCase.EXPECT().
			Update(gomock.Any(), id, input).
			Return(errors.New("update error"))

		body, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPut, "/vehicles/"+id, bytes.NewReader(body))
		req = muxSetURLParam(req, "id", id)
		w := httptest.NewRecorder()

		suite.handler.Update(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		suite.Equal(http.StatusInternalServerError, resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		suite.Contains(string(bodyResp), "update error")
	})
}

func muxSetURLParam(r *http.Request, key, value string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, &chi.Context{
		URLParams: chi.RouteParams{
			Keys:   []string{key},
			Values: []string{value},
		},
	}))
}
