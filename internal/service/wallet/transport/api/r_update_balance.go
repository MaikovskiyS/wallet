package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"wallet/internal/domain"
)

const (
	BalanceIncrease = 1
	BalanceDecrease = 0
)

type UpdateBalanceDTO struct {
	TransactionType string `json:"transaction_type"`
	WalletId        string `json:"wallet_id"`
	Amount          string `json:"amount"`
}
type UpdateBalanceResponce struct {
	Msg string
}

func (a *api) UpdateBalance(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPatch {
		ErrBadRequest.AddLocation("UpdateBalance-ParseMethod")
		ErrBadRequest.SetErr(errors.New("wrong method"))
		return ErrBadRequest
	}

	dto := UpdateBalanceDTO{}
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		ErrBadRequest.AddLocation("UpdateBalance-DecodeBody")
		ErrBadRequest.SetErr(err)
		return ErrBadRequest
	}
	r.Body.Close()

	transaction, err := dto.toModel()
	if err != nil {
		return err
	}
	err = a.u.UpdateBalance(r.Context(), transaction)
	if err != nil {
		return err
	}
	//TODO: add response
	return nil
}

// toModel parse and validate input request data. Return Transaction model
func (d *UpdateBalanceDTO) toModel() (domain.Transaction, error) {
	tType, err := strconv.Atoi(d.TransactionType)
	if err != nil {
		ErrBadRequest.AddLocation("UpdateBalance-ParseTransactionType")
		ErrBadRequest.SetErr(err)
		return domain.Transaction{}, ErrBadRequest
	}
	if tType != BalanceIncrease && tType != BalanceDecrease {
		ErrBadRequest.AddLocation("UpdateBalance-ValidateTransactionType")
		ErrBadRequest.SetErr(errors.New("wrong transaction type"))
		return domain.Transaction{}, ErrBadRequest
	}
	walletId, err := strconv.Atoi(d.WalletId)
	if err != nil {
		ErrBadRequest.AddLocation("UpdateBalance-ParseWalletId")
		ErrBadRequest.SetErr(err)
		return domain.Transaction{}, ErrBadRequest
	}
	if walletId <= 0 {
		ErrBadRequest.AddLocation("UpdateBalance-ValidateWalletId")
		ErrBadRequest.SetErr(errors.New("walletId should be greater than zero"))
		return domain.Transaction{}, ErrBadRequest
	}
	amount, err := strconv.ParseFloat(d.Amount, 64)
	if err != nil {
		ErrBadRequest.AddLocation("UpdateBalance-ParseAmount")
		ErrBadRequest.SetErr(err)
		return domain.Transaction{}, ErrBadRequest
	}
	if walletId <= 0 {
		ErrBadRequest.AddLocation("UpdateBalance-ValidateAmount")
		ErrBadRequest.SetErr(errors.New("amount should be greater than zero"))
		return domain.Transaction{}, ErrBadRequest
	}
	tr := domain.Transaction{
		TransactionType: uint8(tType),
		WalletId:        uint64(walletId),
		Amount:          amount,
	}
	return tr, nil
}
