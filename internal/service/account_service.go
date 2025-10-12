package service

import (
	"context"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/repository"
)

type AccountService struct {
	repository *repository.Queries
}

func NewAccountService(repository *repository.Queries) *AccountService {
	return &AccountService{repository: repository}
}

func (a *AccountService) GetBalance(ctx context.Context, account *domain.Account) *domain.Account {
	transactions := a.repository.GetTransactionFromAccount(account.Id)

	for _, transaction := range transactions {
		if transaction.Type == domain.OrderSell {
			if transaction.Currency == "BRL" {
				account.UpdateBTCBalance(transaction.ConvertToBTC())
			} else {
				account.UpdateBRLBalance(transaction.ConvertToBRL())
			}
		} else {
			if transaction.Currency == "BRL" {
				account.UpdateBRLBalance(transaction.ConvertToBRL())
			} else {
				account.UpdateBTCBalance(transaction.ConvertToBTC())
			}
		}
	}

	return account
}
