package repository

import (
	"context"
	"database/sql"
	"wallet-service/internal/model"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetWalletByUserID(ctx context.Context, userID string) (*model.Wallet, error) {
	var wallet model.Wallet
	query := `SELECT id, balance, version FROM wallets WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&wallet.ID, &wallet.Balance, &wallet.Version)
	if err == sql.ErrNoRows {
		return nil, model.ErrWalletNotFound
	}
	return &wallet, err
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, tx *sql.Tx, walletID string, newBalance float64, currentVersion int) error {
	result, err := tx.ExecContext(ctx,
		`UPDATE wallets SET balance = $1, version = version + 1 
		 WHERE id = $2 AND version = $3`,
		newBalance, walletID, currentVersion,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return model.ErrConcurrentUpdate
	}
	return nil
}

func (r *WalletRepository) CreateTransaction(ctx context.Context, tx *sql.Tx, walletID string, amount float64, txType, status string) error {
	_, err := tx.ExecContext(ctx,
		`INSERT INTO transactions (wallet_id, amount, transaction_type, status) 
		 VALUES ($1, $2, $3, $4)`,
		walletID, amount, txType, status,
	)
	return err
}

func (r *WalletRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}