package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	orderUID := chi.URLParam(r, "uid")
	if orderUID == "" {
		http.Error(w, "missing order uid", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(r.Context(), orderUID)
	if err != nil {
		slog.Error("failed to get order", "order_uid", orderUID, "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if order == nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		slog.Error("failed to encode response", "order_uid", orderUID, "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
