package dto

import "github.com/google/uuid"

type OrderInput struct {
	Type      string `json:"type"`
	Currency  string `json:"currency"`
	Amount    uint64 `json:"amount"`
	AccountId string `json:"accountId"`
}

type OrderOutput struct {
	Id uuid.UUID `json:"id"`
}

func NewOrderOutput(id uuid.UUID) *OrderOutput {
	return &OrderOutput{Id: id}
}
