package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/central-limit-order-book/internal/domain"
	"github.com/Tulkdan/central-limit-order-book/internal/dto"
	"github.com/Tulkdan/central-limit-order-book/internal/service"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func NewAccountHandler(AccountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: AccountService,
	}
}

func (a *AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue("id")

	account := domain.NewAccount(accountId)
	account = a.AccountService.GetBalance(r.Context(), account)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		dto.AccountOutput{
			Id:      account.Id,
			Balance: account.Balance,
		},
	)
}
