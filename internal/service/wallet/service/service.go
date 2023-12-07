package service

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/service/wallet/ports/repository"
)

type Service struct {
	conv  Converter
	store Repository
}

// New wallet service
func New(r Repository, c Converter) *Service {
	return &Service{
		store: r,
		conv:  c,
	}
}

// Save saves wallet into database; Returning Id
func (s *Service) Create(ctx context.Context, w domain.Wallet) (uint64, error) {
	return s.store.Save(ctx, w)
}

// UpdateBalance converting amount to decimal18 and modify wallet balance
func (s *Service) UpdateBalance(ctx context.Context, t domain.Transaction) error {

	evmAmount, err := s.conv.EvmDecimal(t.Amount)
	if err != nil {
		return err
	}

	p := repository.UpdateParams{
		TransactionType: t.TransactionType,
		WalletId:        t.WalletId,
	}
	switch t.TransactionType {
	case 1:
		p.Amount = int64(evmAmount)
	case 0:
		p.Amount = int64(-evmAmount)
	}
	err = s.store.Update(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

// TransferAmount ...
func (s *Service) TransferAmount(ctx context.Context, t domain.Transfer) error {

	evmAmount, err := s.conv.EvmDecimal(t.Amount)
	if err != nil {
		return err
	}

	err = s.store.Update(ctx, repository.UpdateParams{TransactionType: 0, WalletId: t.From, Amount: -int64(evmAmount)})
	if err != nil {
		return err
	}
	err = s.store.Update(ctx, repository.UpdateParams{TransactionType: 1, WalletId: t.To, Amount: int64(evmAmount)})
	if err != nil {
		return err
	}
	return nil
}
