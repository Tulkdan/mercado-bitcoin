package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/central-limit-order-book/internal/dto"
	"github.com/Tulkdan/central-limit-order-book/internal/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: OrderService,
	}
}

func (p *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.OrderInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := p.OrderService.CreateOrder(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
