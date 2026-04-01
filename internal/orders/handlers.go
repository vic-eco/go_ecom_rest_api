package orders

import (
	"github.com/go-chi/chi/v5"
	"github.com/vic-eco/go_ecom_rest_api/internal/errors"
	"github.com/vic-eco/go_ecom_rest_api/internal/json"
	"log/slog"
	"net/http"
	"strconv"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var payload createOrderParams
	if err := json.Read(r, &payload); err != nil {
		slog.Error("error parsing data", "error", err)
		json.WriteError(w, http.StatusBadRequest, "error parsing data")
		return
	}

	if payload.CustomerID == 0 {
		slog.Error("customer id is missing", "error", "BadRequest")
		json.WriteError(w, http.StatusBadRequest, "customer id is missing")
		return
	}
	if len(payload.Items) == 0 {
		slog.Error("at least one item required", "error", "BadRequest")
		json.WriteError(w, http.StatusBadRequest, "at least one item required")
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), payload)
	if err != nil {

		if errors.IsNotFound(err) {
			slog.Error("item in payload not found", "error", "BadRequest")
			json.WriteError(w, http.StatusBadRequest, "item in payload not found")
			return
		}
		if errors.IsNoStock(err) {
			slog.Error("item out of stock", "error", "BadRequest")
			json.WriteError(w, http.StatusBadRequest, "item out of stock")
			return
		}

		slog.Error("error placing order", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "error placing order")
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)

}

func (h *handler) FindOrderByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("invalid order id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, "invalid order id")
		return
	}

	order, err := h.service.FindOrderByID(r.Context(), id)
	if err != nil {
		if errors.IsNotFound(err) {
			slog.Error("order not found", "error", "NotFound")
			json.WriteError(w, http.StatusNotFound, "order not found")
			return
		}
		slog.Error("error fetching order", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "error fetching order")
		return
	}

	json.Write(w, http.StatusOK, order)

}
