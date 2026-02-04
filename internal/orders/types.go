package orders

import (
	"context"

	repo "github.com/Jakob-Kaae/Go.Demo/internal/adapters/postgresql/sqlc"
)

type OrderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type CreateOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []OrderItem `json:"items"`
}

type Service interface {
	CreateOrder(ctx context.Context, tempOrder CreateOrderParams) (repo.Order, error)
	GetOrders(ctx context.Context) ([]repo.Order, error)
}
