package service

// type Params struct {
// 	From   int64
// 	To     int64
// 	Amount int64
// }

type GetParams struct {
	from   int64
	to     int64
	amount int64
}

func (p *GetParams) From() int64 {
	return p.from
}
func (p *GetParams) To() int64 {
	return p.to
}
func (p *GetParams) Amount() int64 {
	return p.amount
}
