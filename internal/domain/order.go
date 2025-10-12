package domain

import (
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"  // when it needs to send payment to provider
	StatusApproved Status = "approved" // when provider successfully charged
	StatusRejected Status = "canceled" // when provider rejected payment
)

type OrderType string

const (
	OrderSell OrderType = "sell"
	OrderBuy  OrderType = "buy"
)

type Order struct {
	Id        uuid.UUID
	Amount    uint64
	Currency  string
	Type      OrderType
	AccountId string
	Status    Status
	CreatedAt time.Time

	mu sync.Mutex
}

func NewOrder(amount uint64, currency, accountId, orderType string) (*Order, error) {
	isoFormatRgx := regexp.MustCompile(`^[A-Z]{3}$`)
	if !isoFormatRgx.Match([]byte(currency)) {
		return nil, errors.New("Invalid Currency")
	}

	if amount == 0 {
		return nil, errors.New("Amount is required")
	}

	if strings.TrimSpace(accountId) == "" {
		return nil, errors.New("AccountId is required")
	}

	if strings.TrimSpace(orderType) == "" {
		return nil, errors.New("OrderType is required")
	}

	if orderType != string(OrderSell) && orderType != string(OrderBuy) {
		return nil, errors.New("Invalid OrderType")
	}

	typeEnum := OrderBuy
	if orderType == string(OrderSell) {
		typeEnum = OrderSell
	}

	return &Order{
		Id:        uuid.New(),
		Amount:    amount,
		Currency:  currency,
		AccountId: accountId,
		Type:      typeEnum,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}, nil
}

func (p *Order) UpdateStatus(status Status) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Status = status
}

func (p *Order) ConvertToBTC() uint64 {
	if p.Currency == "BTC" {
		return p.Amount
	}
	return p.Amount / 1_000
}

func (p *Order) ConvertToBRL() uint64 {
	if p.Currency == "BRL" {
		return p.Amount
	}
	return p.Amount * 1_000
}
