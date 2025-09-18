package http

import (
	"encoding/json"
	"net/http"

	"github.com/NicolasNSC/catalog-service-fiap/internal/usecase"
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
	var input usecase.InputCreateVehicleDTO
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
