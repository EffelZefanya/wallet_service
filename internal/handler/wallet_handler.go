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
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBalance, err := h.service.Withdraw(c.Request.Context(), req.UserID, req.Amount)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, WithdrawResponse{
		Message: "withdrawal successful",
		RemainingBalance: newBalance,
	})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	var req BalanceRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := h.service.GetBalance(c.Request.Context(), req.UserID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{
		UserID:  req.UserID,
		Balance: balance,
	})
}

func (h *WalletHandler) handleServiceError(c *gin.Context, err error) {
	code := http.StatusInternalServerError
	msg := "internal server error"

	if errors.Is(err, model.ErrInsufficientFunds) {
		code = http.StatusUnprocessableEntity
		msg = err.Error()
	} else if errors.Is(err, model.ErrConcurrentUpdate) {
		code = http.StatusConflict
		msg = err.Error()
	} else if errors.Is(err, model.ErrWalletNotFound) {
		code = http.StatusNotFound
		msg = err.Error()
	}
	c.JSON(code, gin.H{"error": msg})
}