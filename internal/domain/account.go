package domain

import "sync"

type Account struct {
	Id         string
	BRLBalance int64
	BTCBalance int64

	mu sync.Mutex
}

func NewAccount(id string) *Account {
	return &Account{Id: id, BRLBalance: 0, BTCBalance: 0}
}

func (a *Account) UpdateBTCBalance(value uint64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.BTCBalance += int64(value)
}

func (a *Account) UpdateBRLBalance(value uint64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.BRLBalance += int64(value)
}
