package service

import (
	"fmt"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/repository"
)

type SellService struct {
	repository *repository.Queries
}

func NewSellService(repository *repository.Queries) *SellService {
	return &SellService{repository: repository}
}

func (s *SellService) MakeSales() {
	allTransactions := s.repository.GetAllPendingTransactions()

	for _, transaction := range allTransactions {
		fmt.Printf("%+v\n", transaction)

		for _, transaction2 := range allTransactions {
			if transaction.Id == transaction2.Id {
				continue
			}

			if transaction.Type == transaction2.Type {
				continue
			}

			if transaction.ConvertAmount() == transaction2.ConvertAmount() {
				transaction.UpdateStatus(domain.StatusApproved)
				transaction2.UpdateStatus(domain.StatusApproved)
			}
		}
	}
}
