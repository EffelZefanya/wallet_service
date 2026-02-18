package service

import (
	"context"
	"database/sql"
	"errors"
	"wallet-service/internal/model"
)

// WalletRepository defines the interface for wallet data access.
// This allows for mocking in tests and decouples the service from the concrete repository.
type WalletRepository interface {
	GetWallet(ctx context.Context, userID string) (float64, int, string, error)
	UpdateBalance(ctx context.Context, tx *sql.Tx, walletID string, newBalance float64, currentVersion int) error
	CreateTransaction(ctx context.Context, tx *sql.Tx, walletID string, amount float64, txType, status string) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type WalletService struct {
	repo WalletRepository
	maxRetries int
}

func NewWalletService(repo WalletRepository, maxRetries int) *WalletService {
	if maxRetries <= 0 {
		maxRetries = 3
	}
	return &WalletService{repo: repo, maxRetries: maxRetries}
}

func (s *WalletService) GetBalance(ctx context.Context, userID string) (float64, error) {
	balance, _, _, err := s.repo.GetWallet(ctx, userID)
	return balance, err
}

func (s *WalletService) Withdraw(ctx context.Context, userID string, amount float64) (float64, error) {
	for i := 0; i < s.maxRetries; i++ {
		balance, version, walletID, err := s.repo.GetWallet(ctx, userID)
		if err != nil {
			return 0, err
		}

		if balance < amount {
			return 0, model.ErrInsufficientFunds
		}

		tx, err := s.repo.BeginTx(ctx)
		if err != nil {
			return 0, err
		}

		// Use a closure to handle the transaction scope and defer rollback safely within the loop
		newBalance, err := func() (float64, error) {
			defer tx.Rollback()

			newBalance := balance - amount
			if err := s.repo.UpdateBalance(ctx, tx, walletID, newBalance, version); err != nil {
				return 0, err
			}

			if err := s.repo.CreateTransaction(ctx, tx, walletID, amount, "WITHDRAWAL", "SUCCESS"); err != nil {
				return 0, err
			}

			return newBalance, tx.Commit()
		}()

		if err != nil {
			if errors.Is(err, model.ErrConcurrentUpdate) {
				continue // Retry the entire process (read + write)
			}
			return 0, err
		}

		return newBalance, nil
	}

	return 0, model.ErrConcurrentUpdate
}