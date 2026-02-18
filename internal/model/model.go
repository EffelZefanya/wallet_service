package service

import "time"

type Wallet struct {
	ID        string
	UserID    string
	Balance   float64
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionRecord struct {
	ID        string
	Amount    float64
	Type      string
	Status    string
	CreatedAt time.Time
}