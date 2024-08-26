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
	Addresses []*SolAddress `json:"addresses"`
}

// POST /address/name
type NameAddressReq struct {
	SolAddress string `json:"sol_address"`
	Nickname   string `json:"nickname"`
}

// POST /address/delete
type DeleteAddressReq struct {
	SolAddress string `json:"sol_address"`
}
