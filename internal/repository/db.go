package repository

import (
	"context"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
)

type Queries struct {
	transactions []*domain.Order
}

func New() *Queries {
	return &Queries{
		transactions: make([]*domain.Order, 0),
	}
}

func (q *Queries) SaveTransaction(ctx context.Context, order *domain.Order) {
	q.transactions = append(q.transactions, order)
}
