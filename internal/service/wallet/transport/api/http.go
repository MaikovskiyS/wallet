package api

import (
	"context"
	"wallet/internal/apperrors"
	"wallet/internal/domain"
)

const location = "Wallet-Api-"

var (
	ErrBadRequest = apperrors.New(apperrors.ErrBadRequest, location)
)

type Wallet interface {
	Create(ctx context.Context, w domain.Wallet) (uint64, error)
	UpdateBalance(ctx context.Context, t domain.Transaction) error
	TransferAmount(ctx context.Context, t domain.Transfer) error
}
type api struct {
	u Wallet
}

func New(u Wallet) *api {
	return &api{
		u: u,
	}
}
