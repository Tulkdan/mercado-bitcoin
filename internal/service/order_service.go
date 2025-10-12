package service

import (
	"context"
	"errors"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/dto"
	"github.com/Tulkdan/central-limit-order-book/internal/repository"
	"github.com/google/uuid"
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

	p.repository.SaveTransaction(order)

	return dto.NewOrderOutput(order.Id), nil
}

func (p *OrderService) CancelOrder(ctx context.Context, orderId uuid.UUID) error {
	order := p.repository.GetTransaction(orderId)

	if order == nil {
		return errors.New("Not Found")
	}

	if order.Status == domain.StatusApproved {
		return errors.New("Can't cancel an approved transaction")
	}

	order.UpdateStatus(domain.StatusRejected)

	return nil
}
