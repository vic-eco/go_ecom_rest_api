package products

import (
	"context"
	"database/sql"
	"errors"
	myErr "github.com/vic-eco/go_ecom_rest_api/internal/errors"
	repo "github.com/vic-eco/go_ecom_rest_api/internal/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, params createProductParams) (repo.Product, error)
}
type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repo.Product{}, myErr.ErrNotFound
		}
		return repo.Product{}, err
	}
	return product, nil
}

func (s *svc) CreateProduct(ctx context.Context, params createProductParams) (repo.Product, error) {
	product, err := s.repo.CreateProduct(ctx, repo.CreateProductParams{
		Name:         params.Name,
		Quantity:     params.Quantity,
		PriceInCents: params.PriceInCents,
	})
	if err != nil {
		return repo.Product{}, err
	}
	return product, nil
}
