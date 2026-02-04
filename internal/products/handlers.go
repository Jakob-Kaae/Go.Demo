package products

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Jakob-Kaae/Go.Demo/internal/json"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.ListProducts(r.Context())
	if err != nil {
		slog.Error("Failed to list products", "error", err)
		http.Error(w, "Failed to list products", http.StatusInternalServerError)
		return
	}
	json.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	// extract id from the URL path (chi router provides URLParam)
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "missing id in url", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Invalid id in URL", "id", idStr, "error", err)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	slog.Info("GetProductById called", "id", id)
	product, err := h.Service.GetProductById(r.Context(), id)
	if err != nil {
		slog.Error("Failed to get product by ID", "error", err)
		http.Error(w, "Failed to get product by ID", http.StatusInternalServerError)
		return
	}
	json.WriteJSON(w, http.StatusOK, product)
}
