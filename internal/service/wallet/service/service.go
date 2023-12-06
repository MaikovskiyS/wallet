package service

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/service/wallet/ports/repository"
)

type service struct {
	conv  Converter
	store Repository
}

// New wallet service
func New(r Repository, c Converter) *service {
	return &service{
		store: r,
		conv:  c,
	}
}

// Save saves wallet into database; Returning Id
func (s *service) Create(ctx context.Context, w domain.Wallet) (uint64, error) {
	return s.store.Save(ctx, w)
}

/*
Ожидаемый результат: в зависимости от типа транзакции, должен
измениться баланс указанного кошелька (либо увеличиться, либо
уменьшиться соответственно). Транзакция должна быть записана в
соответствующую таблицу. Помимо входных данных, запись в
таблице должна содержать поле updated_balance с балансом
кошелька после выполнения транзакции.

1. определить тип транзакции
2. создать параметры запроса к бд update balance params; amount зависит от типа транзакции

3. создаем новую транзакцию
*/

/*

 */
// UpdateBalance converting amount to decimal18 and modify wallet balance
func (s *service) UpdateBalance(ctx context.Context, t domain.Transaction) error {

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
func (s *service) TransferAmount(ctx context.Context, t domain.Transfer) error {

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
