package domain

type Wallet struct {
	Id           uint64
	CurrencyName string
	Balance      float64
}

type Transaction struct {
	TransactionType uint8
	WalletId        uint64
	Amount          float64
}

type Transfer struct {
	From   uint64
	To     uint64
	Amount float64
}
