package repository

import "wallet/internal/domain"

type walletRow struct {
	id      int64  `db:"id"`
	name    string `db:"name"`
	balance uint64 `db:"balance"`
}

func (row *walletRow) toModel() domain.Wallet {
	w := domain.Wallet{
		Id:           uint64(row.id),
		CurrencyName: row.name,
		Balance:      float64(row.balance),
	}
	return w
}
