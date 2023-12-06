package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"wallet/internal/domain"
)

type TransferAmountDTO struct {
	From   string `json:"from_wallet_id"`
	To     string `json:"to_wallet_id"`
	Amount string `json:"amount"`
}

// TransferAmount handle input Json, parse data to domain entity, call service layer method, send responce to client
func (a *api) TransferAmount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPatch {
		ErrBadRequest.AddLocation("UpdateBalance-ParseMethod")
		ErrBadRequest.SetErr(errors.New("wrong method"))
		return ErrBadRequest
	}

	dto := TransferAmountDTO{}
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		ErrBadRequest.AddLocation("UpdateBalance-DecodeBody")
		ErrBadRequest.SetErr(err)
		return ErrBadRequest
	}
	r.Body.Close()

	transfer, err := dto.toModel()
	if err != nil {
		return err
	}
	err = a.u.TransferAmount(r.Context(), transfer)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("transfer success"))
	return nil
}

// toModel validate input data and mapping to domain entity
func (d *TransferAmountDTO) toModel() (domain.Transfer, error) {

	fromWalletId, err := strconv.ParseUint(d.From, 10, 64)
	if err != nil {
		ErrBadRequest.AddLocation("toModel-ParseFieldFrom")
		ErrBadRequest.SetErr(err)
		return domain.Transfer{}, ErrBadRequest
	}

	toWalletId, err := strconv.ParseUint(d.To, 10, 64)
	if err != nil {
		ErrBadRequest.AddLocation("toModel-ParseFieldTo")
		ErrBadRequest.SetErr(err)
		return domain.Transfer{}, ErrBadRequest
	}

	amount, err := strconv.ParseFloat(d.Amount, 64)
	if err != nil {
		ErrBadRequest.AddLocation("toModel-ParseFieldTo")
		ErrBadRequest.SetErr(err)
		return domain.Transfer{}, ErrBadRequest
	}

	t := domain.Transfer{From: fromWalletId, To: toWalletId, Amount: amount}

	return t, nil
}
