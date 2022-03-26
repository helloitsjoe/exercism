package account

import "sync"

type Account struct {
	balance int64
	isOpen  bool
	mu      sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	a := Account{balance: amount, isOpen: true}
	return &a
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	a.isOpen = false
	return a.balance, true
}
