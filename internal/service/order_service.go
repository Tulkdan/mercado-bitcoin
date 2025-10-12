package service

import (
	"context"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/dto"
	"github.com/Tulkdan/central-limit-order-book/internal/repository"
)

type OrderService struct {
	repository *repository.Queries
}

func NewOrderService(repository *repository.Queries) *OrderService {
	return &OrderService{repository: repository}
}

func (p *OrderService) CreateOrder(ctx context.Context, input dto.OrderInput) (*dto.OrderOutput, error) {
	order, err := domain.NewOrder(input.Amount, input.Currency, input.AccountId, input.Type)
	if err != nil {
		return nil, err
	}

	p.repository.SaveTransaction(ctx, order)

	return dto.NewOrderOutput(order.Id), nil
}
