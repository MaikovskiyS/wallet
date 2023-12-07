package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"wallet/internal/domain"
)

type SaveWalletRequest struct {
	Coin string `json:"coin"`
}

type SaveWalletResponce struct {
	Msg string `json:"msg"`
	Id  uint64 `json:"id"`
}

// CreateWallet  handle input Json, parse data to domain entity, call service layer method, send responce to client
func (a *api) CreateWallet(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		ErrBadRequest.AddLocation("Save-ParseMethod")
		ErrBadRequest.SetErr(errors.New("wrong method"))
		return ErrBadRequest
	}

	dto := SaveWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		ErrBadRequest.AddLocation("Save-DecodeBody")
		ErrBadRequest.SetErr(err)
		return ErrBadRequest
	}
	r.Body.Close()

	wallet, err := dto.toModel()
	if err != nil {
		return err
	}
	id, err := a.u.Create(r.Context(), wallet)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := SaveWalletResponce{Id: id, Msg: "wallet created"}
	err = resp.Send(w)
	if err != nil {
		return err
	}
	return nil
}

func (r *SaveWalletRequest) toModel() (domain.Wallet, error) {
	if r.Coin == "" {
		ErrBadRequest.AddLocation("ToModel-ParseCoinName")
		ErrBadRequest.SetErr(errors.New("coinName required"))
		return domain.Wallet{}, ErrBadRequest
	}
	wallet := domain.Wallet{CurrencyName: r.Coin, Balance: 0}

	return wallet, nil
}

// toModel validate input data and mapping to domain entity
func (resp *SaveWalletResponce) Send(w http.ResponseWriter) error {
	idBytes, err := json.Marshal(&resp)
	if err != nil {
		return err
	}
	w.Write(idBytes)
	return nil
}
