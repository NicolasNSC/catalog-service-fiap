package http

import (
	"encoding/json"
	"net/http"

	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
	"github.com/NicolasNSC/catalog-service-fiap/internal/usecase"
	"github.com/go-chi/chi"
)

type VehicleHandler struct {
	useCase usecase.VehicleUseCaseInterface
}

func NewVehicleHandler(useCase usecase.VehicleUseCaseInterface) *VehicleHandler {
	return &VehicleHandler{
		useCase: useCase,
	}
}

// Create is the handler for the POST /vehicles endpoint.
// @Summary      Create a new vehicle
// @Description  Adds a new vehicle to the catalog
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        vehicle  body      dto.InputCreateVehicleDTO  true  "Vehicle data to create"
// @Success      201      {object}  dto.OutputCreateVehicleDTO
// @Failure      400      {string}  string "Invalid request body"
// @Failure      500      {string}  string "Failed to create vehicle"
// @Router       /vehicles/add [post]
func (h *VehicleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.InputCreateVehicleDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	output, err := h.useCase.Create(r.Context(), input)
	if err != nil {
		http.Error(w, "Failed to create vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Update is the handler for the PUT /vehicles/{id} endpoint.
// @Summary      Update an existing vehicle
// @Description  Updates the data of a vehicle identified by its ID
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        id       path      string                     true  "Vehicle ID"
// @Param        vehicle  body      dto.InputUpdateVehicleDTO  true  "Vehicle data to update"
// @Success      200      {string}  string "OK"
// @Failure      400      {string}  string "Invalid request body or ID"
// @Failure      404      {string}  string "Vehicle not found"
// @Failure      500      {string}  string "Internal server error"
// @Router       /vehicles/{id} [put]
func (h *VehicleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "VehicleID is required", http.StatusBadRequest)
		return
	}

	var input dto.InputUpdateVehicleDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.useCase.Update(r.Context(), id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
