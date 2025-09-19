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
