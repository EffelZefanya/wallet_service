package model

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrWalletNotFound    = errors.New("wallet not found")
	ErrConcurrentUpdate  = errors.New("concurrent update error, please try again")
)