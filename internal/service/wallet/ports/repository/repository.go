package repository

import (
	"context"
	"database/sql"
	"errors"
	"wallet/internal/apperrors"
	"wallet/internal/domain"

	"github.com/jmoiron/sqlx"
)

const location = "Wallet-Repository-"

var (
	ErrInternal = apperrors.New(apperrors.ErrInternal, location)
	ErrNotFound = apperrors.New(apperrors.ErrNotFound, location)
)

type repo struct {
	cl *sqlx.DB
}

// Repository constructor
func New(cl *sqlx.DB) *repo {
	return &repo{
		cl: cl,
	}
}

// Save saves wallet entity into database and return id
func (r *repo) Save(ctx context.Context, w domain.Wallet) (uint64, error) {
	query := "INSERT INTO wallets(name, balance) VALUES($1,$2) RETURNING ID"
	var id uint64

	err := r.cl.QueryRowContext(ctx, query, w.CurrencyName, w.Balance).Scan(&id)
	if err != nil {
		ErrInternal.AddLocation("Save-QueryRowContext")
		ErrInternal.SetErr(err)
		return 0, ErrInternal
	}
	return id, nil
}

// GetById getting wallet entity by Id
func (r *repo) GetById(ctx context.Context, id uint64) (domain.Wallet, error) {

	query := "SELECT id, name, balance FROM wallets WHERE id=$1"
	wRow := walletRow{}
	err := r.cl.QueryRowContext(ctx, query, id).Scan(&wRow.id, &wRow.name, &wRow.balance)
	if err != nil {
		if err == sql.ErrNoRows {
			ErrNotFound.AddLocation("GetById-QueryRowContext")
			ErrNotFound.SetErr(errors.New("wallet not found"))
			return domain.Wallet{}, ErrNotFound
		}
		ErrInternal.AddLocation("GetById-QueryRowContext")
		ErrInternal.SetErr(err)
		return domain.Wallet{}, ErrInternal
	}

	wallet := wRow.toModel()
	return wallet, nil

}

type UpdateParams struct {
	TransactionType uint8
	WalletId        uint64
	Amount          int64
}

// Update balance by wallet_id and creating transacion_row
func (r *repo) Update(ctx context.Context, p UpdateParams) error {
	tx, err := r.cl.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		ErrInternal.AddLocation("Update-BeginTx")
		ErrInternal.SetErr(err)
		return ErrInternal
	}
	defer func() error {
		if err := tx.Rollback(); err != nil {
			ErrInternal.AddLocation("Update-Rollback")
			ErrInternal.SetErr(err)
			return ErrInternal

		}
		return errors.New("tx rollback")
	}()

	query := "UPDATE wallets SET balance=balance+$1 WHERE id=$2 RETURNING balance"
	var updatedBalance int64
	err = tx.QueryRowContext(ctx, query, p.Amount, p.WalletId).Scan(&updatedBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			ErrNotFound.AddLocation("Update-QueryRowContext")
			ErrNotFound.SetErr(errors.New("wallet not found"))
			return ErrNotFound
		}
		ErrInternal.AddLocation("Update-QueryRowContext")
		ErrInternal.SetErr(err)
		return ErrInternal
	}

	query = "INSERT INTO transactions (transaction_type, wallet_id, amount, updated_balance)VALUES($1,$2,$3,$4)"
	_, err = tx.ExecContext(ctx, query, p.TransactionType, p.WalletId, p.Amount, updatedBalance)
	if err != nil {
		ErrInternal.AddLocation("Update-ExecContext")
		ErrInternal.SetErr(err)
		return ErrInternal
	}
	err = tx.Commit()
	if err != nil {
		ErrInternal.AddLocation("Update-Commit")
		ErrInternal.SetErr(err)
		return ErrInternal
	}
	return nil
}

type TransferParams struct {
	From   uint64
	To     uint64
	Amount float64
}
