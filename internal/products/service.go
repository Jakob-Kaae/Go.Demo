package products

import (
	"context"

	repo "github.com/Jakob-Kaae/Go.Demo/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(r repo.Querier) Service {
	return &svc{
		repo: r,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.GetProductById(ctx, id)
}
