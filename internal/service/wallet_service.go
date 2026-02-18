package service

import (
	"errors"
	"wallet-microservice/internal/repository"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) GetBalance(userID string) (float64, error) {
	balance, _, _, err := s.repo.GetWallet(userID)
	return balance, err
}

func (s *WalletService) Withdraw(userID string, amount float64) error {
	balance, version, walletID, err := s.repo.GetWallet(userID)
	if err != nil {
		return err
	}

	if balance < amount {
		return errors.New("insufficient funds")
	}

	tx, err := s.repo.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	newBalance := balance - amount
	if err := s.repo.UpdateBalance(tx, walletID, newBalance, version); err != nil {
		return err
	}

	if err := s.repo.CreateTransaction(tx, walletID, amount, "WITHDRAWAL", "SUCCESS"); err != nil {
		return err
	}

	return tx.Commit()
}