package handler

import (
	"encoding/json"
	"net/http"
	"Taxi/internal/models"
	"Taxi/internal/service"
)

type RideHandler struct{
	priceService *service.PriceService
}

func NewRideHandler (svc *service.PriceService) *RideHandler{
	return &RideHandler{
		priceService: svc,
	}
}

func (h *RideHandler) CalculatePrice (w http.ResponseWriter, r *http.Request){
	var req models.RideRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	rez := h.priceService.CalculatePrice(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rez)
}