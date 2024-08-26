package handler

// POST /address/submit
type SubmitAddressReq struct {
	SolAddress string `json:"sol_address"`
}

type SubmitAddressRes struct {
	NewAddress bool `json:"new_address"`
}

// GET /address
type GetAddressesRes struct {
	Addresses    []*SolAddress `json:"addresses"`
	ExchangeRate float64       `json:"exchange_rate"`
}

// POST /address/name
type NameAddressReq struct {
	SolAddress string `json:"sol_address"`
	Nickname   string `json:"nickname"`
}
