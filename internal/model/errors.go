package model

import "errors"

var (
	ErrConcurrentUpdate  = errors.New("concurrent update detected")
	ErrWalletNotFound    = errors.New("wallet not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)