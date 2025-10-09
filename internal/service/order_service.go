package service

import (
	"context"
	"fmt"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/dto"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (p *OrderService) CreateOrder(ctx context.Context, input dto.OrderInput) (*dto.OrderOutput, error) {
	order, err := domain.NewOrder(input.Amount, input.Currency, input.AccountId, input.Type)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", order)

	// TODO: save to db

	return dto.NewOrderOutput(order.Id), nil
}
