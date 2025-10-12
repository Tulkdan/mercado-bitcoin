package domain

import "sync"

type Account struct {
	Id      string
	Balance int64

	mu sync.Mutex
}

func NewAccount(id string) *Account {
	return &Account{Id: id, Balance: 0}
}

func (a *Account) UpdateBalance(value uint64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += int64(value)
}
