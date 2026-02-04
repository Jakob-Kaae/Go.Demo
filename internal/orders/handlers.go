package orders

import (
	"log"
	"net/http"

	"github.com/Jakob-Kaae/Go.Demo/internal/json"
)

type Handler struct {
	Service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating an order goes here
	var tempOrder CreateOrderParams
	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.Service.CreateOrder(r.Context(), tempOrder)
	if err != nil {
		log.Println(err)

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.WriteJSON(w, http.StatusCreated, createdOrder)
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting orders goes here
	orders, err := h.Service.GetOrders(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	json.WriteJSON(w, http.StatusOK, orders)
}
