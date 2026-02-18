package service

import (
	"context"
	"database/sql"
	"errors"
	"wallet-service/internal/model"
)
type WalletRepository interface {
	GetWalletByUserID(ctx context.Context, userID string) (*model.Wallet, error)
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
	wallet, err := s.repo.GetWalletByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (s *WalletService) Withdraw(ctx context.Context, userID string, amount float64) (float64, error) {
	for i := 0; i < s.maxRetries; i++ {
		wallet, err := s.repo.GetWalletByUserID(ctx, userID)
		if err != nil {
			return 0, err
		}

		if wallet.Balance < amount {
			return 0, model.ErrInsufficientFunds
		}

		tx, err := s.repo.BeginTx(ctx)
		if err != nil {
			return 0, err
		}

		newBalance, err := func() (float64, error) {
			defer tx.Rollback()

			newBalance := wallet.Balance - amount
			if err := s.repo.UpdateBalance(ctx, tx, wallet.ID, newBalance, wallet.Version); err != nil {
				return 0, err
			}

			if err := s.repo.CreateTransaction(ctx, tx, wallet.ID, amount, string(model.TxTypeWithdrawal), string(model.TxStatusSuccess)); err != nil {
				return 0, err
			}

			return newBalance, tx.Commit()
		}()

		if err != nil {
			if errors.Is(err, model.ErrConcurrentUpdate) {
				continue
			}
			return 0, err
		}

		return newBalance, nil
	}

	return 0, model.ErrConcurrentUpdate
}