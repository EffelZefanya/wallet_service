package handler

type WithdrawRequest struct {
	UserID string  `json:"user_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type WithdrawResponse struct {
	Message          string  `json:"message"`
	RemainingBalance float64 `json:"remaining_balance"`
}

type BalanceResponse struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}