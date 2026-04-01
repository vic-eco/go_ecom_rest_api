package orders

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type OrderProduct struct {
	Name       string `json:"name"`
	Quantity   int32  `json:"quantity"`
	PriceCents int32  `json:"price_cents"`
}

type OrderResponse struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Products  []OrderProduct     `json:"products"`
}
