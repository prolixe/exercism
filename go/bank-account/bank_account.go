package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	isOpen  bool
	m       *sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, true, &sync.Mutex{}}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if !a.isOpen {
		return
	}
	a.isOpen = false
	return a.balance, true
}
func (a *Account) Balance() (balance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if !a.isOpen {
		return
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if !a.isOpen {
		return
	}
	if a.balance+amount < 0 {
		return
	}
	a.balance += amount
	return a.balance, true
}
