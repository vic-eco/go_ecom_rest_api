package products

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

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("error listing products", "error", err)
		json.WriteError(w, http.StatusInternalServerError, "error listing products")
		return
	}
	json.Write(w, http.StatusOK, products)
}

func (h *handler) FindProductByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("invalid product id", "id", idStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.service.FindProductByID(r.Context(), int64(id))
	if err != nil {
		if errors.IsNotFound(err) {
			slog.Error("product not found", "error", err)
			json.WriteError(w, http.StatusNotFound, "product not found")
			return
		}
		slog.Error("error finding product", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload createProductParams
	if err := json.Read(r, &payload); err != nil {
		slog.Error("error parsing data", "error", err)
		json.WriteError(w, http.StatusBadRequest, "error parsing data")
		return
	}

	if payload.PriceInCents <= 0 {
		slog.Error("price is negative or missing", "error", "BadRequest")
		json.WriteError(w, http.StatusBadRequest, "price is negative or missing")
		return
	}
	if payload.Quantity <= 0 {
		slog.Error("quantity is negative or missing", "error", "BadRequest")
		json.WriteError(w, http.StatusBadRequest, "quantity is negative or missing")
		return
	}

	product, err := h.service.CreateProduct(r.Context(), payload)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, "failed to create product")
		return
	}

	json.Write(w, http.StatusCreated, product)
}
