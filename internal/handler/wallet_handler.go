package handler

import (
	"errors"
	"net/http"
	"wallet-service/internal/model"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(s *service.WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}

func (h *WalletHandler) Withdraw(c *gin.Context) {
	// Using our DTO
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	newBalance, err := h.service.Withdraw(c.Request.Context(), req.UserID, req.Amount)
	if err != nil {
		// Professional tip: map specific errors to status codes
		code := http.StatusInternalServerError
		if errors.Is(err, model.ErrInsufficientFunds) {
			code = http.StatusUnprocessableEntity // 422 is great for business rule violations
		} else if errors.Is(err, model.ErrConcurrentUpdate) {
			code = http.StatusConflict // 409 for optimistic locking failures
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, WithdrawResponse{
		Message: "withdrawal successful",
		RemainingBalance: newBalance,
	})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	userID := c.Query("user_id") // Gin makes query params easy
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	balance, err := h.service.GetBalance(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, model.ErrWalletNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{
		UserID:  userID,
		Balance: balance,
	})
}