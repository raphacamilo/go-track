package account

import "sync"

type Account struct {
	mu      *sync.Mutex
	balance int64
	closed  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{&sync.Mutex{}, amount, false}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance, !a.closed
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if newBalance := a.balance + amount; !a.closed && newBalance >= 0 {
		a.balance = newBalance
		return a.balance, true
	}

	return 0, false
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	oldBalance := a.balance
	a.balance = 0

	return oldBalance, a.closed
}
