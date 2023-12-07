package service

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/service/wallet/ports/repository"
)

type Repository interface {
	GetById(ctx context.Context, id uint64) (domain.Wallet, error)
	Save(ctx context.Context, w domain.Wallet) (uint64, error)
	Update(ctx context.Context, p repository.UpdateParams) error
	ExecTx(ctx context.Context, fn func() error) error
}

type Converter interface {
	EvmDecimal(amount float64) (int64, error)
}
