package orders

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	myErr "github.com/vic-eco/go_ecom_rest_api/internal/errors"
	repo "github.com/vic-eco/go_ecom_rest_api/internal/postgresql/sqlc"
)

type Service interface {
	PlaceOrder(ctx context.Context, order createOrderParams) (repo.Order, error)
	FindOrderByID(ctx context.Context, id int64) (OrderResponse, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, myErr.ErrNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, myErr.ErrNoStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}

		_, err = qtx.UpdateProduct(ctx, repo.UpdateProductParams{
			ID:       item.ProductID,
			Quantity: pgtype.Int4{Int32: product.Quantity - item.Quantity, Valid: true},
		})
		if err != nil {
			return repo.Order{}, err
		}
	}

	tx.Commit(ctx)

	return order, nil
}

func (s *svc) FindOrderByID(ctx context.Context, id int64) (OrderResponse, error) {
	order, err := s.repo.FindOrderByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return OrderResponse{}, myErr.ErrNotFound
		}
		return OrderResponse{}, err
	}

	orderItems, err := s.repo.GetOrderItemsByOrderID(ctx, order.ID)
	if err != nil {
		return OrderResponse{}, err
	}

	var response OrderResponse
	response.ID = order.ID
	response.CreatedAt = order.CreatedAt
	response.Products = make([]OrderProduct, 0, len(orderItems))
	for _, item := range orderItems {
		response.Products = append(response.Products, OrderProduct{
			Name:       item.ProductName,
			Quantity:   item.Quantity,
			PriceCents: item.PriceCents,
		})
	}

	return response, nil
}
