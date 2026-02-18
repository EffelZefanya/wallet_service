package handler

type WithdrawRequest struct {
	UserID string  `json:"user_id" binding:"required,uuid"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type WithdrawResponse struct {
	Message          string  `json:"message"`
	RemainingBalance float64 `json:"remaining_balance"`
}

type BalanceRequest struct {
	UserID string `form:"user_id" binding:"required,uuid"`
}

type BalanceResponse struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}