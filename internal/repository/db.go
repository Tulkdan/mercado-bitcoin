package repository

import (
	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/google/uuid"
)

type Queries struct {
	transactions []*domain.Order
}

func New() *Queries {
	return &Queries{
		transactions: make([]*domain.Order, 0),
	}
}

func (q *Queries) SaveTransaction(order *domain.Order) {
	q.transactions = append(q.transactions, order)
}

func (q *Queries) GetTransaction(orderId uuid.UUID) *domain.Order {
	for _, order := range q.transactions {
		if order.Id == orderId {
			return order
		}
	}
	return nil
}

func (q *Queries) GetTransactionFromAccount(accountId string) []*domain.Order {
	orders := []*domain.Order{}

	for _, order := range q.transactions {
		if order.AccountId == accountId && order.Status == domain.StatusApproved {
			orders = append(orders, order)
		}
	}

	return orders
}
