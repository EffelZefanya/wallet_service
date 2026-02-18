package repository

import (
	"database/sql"
	"errors"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetWallet(userID string) (float64, int, string, error) {
	var balance float64
	var version int
	var walletID string
	
	query := `SELECT id, balance, version FROM wallets WHERE user_id = $1`
	err := r.db.QueryRow(query, userID).Scan(&walletID, &balance, &version)
	return balance, version, walletID, err
}

func (r *WalletRepository) UpdateBalance(tx *sql.Tx, walletID string, newBalance float64, currentVersion int) error {
	result, err := tx.Exec(
		`UPDATE wallets SET balance = $1, version = version + 1 
		 WHERE id = $2 AND version = $3`,
		newBalance, walletID, currentVersion,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("concurrent update detected")
	}
	return nil
}

func (r *WalletRepository) CreateTransaction(tx *sql.Tx, walletID string, amount float64, txType, status string) error {
	_, err := tx.Exec(
		`INSERT INTO transactions (wallet_id, amount, transaction_type, status) 
		 VALUES ($1, $2, $3, $4)`,
		walletID, amount, txType, status,
	)
	return err
}

func (r *WalletRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}