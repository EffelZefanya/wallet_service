package model

type TxType string
type TxStatus string

const (
	TxTypeWithdrawal TxType = "WITHDRAWAL"
	TxTypeDeposit    TxType = "DEPOSIT"
)

const (
	TxStatusPending TxStatus = "PENDING"
	TxStatusSuccess TxStatus = "SUCCESS"
	TxStatusFailed  TxStatus = "FAILED"
)