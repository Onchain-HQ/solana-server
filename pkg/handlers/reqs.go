package handler

// POST /address
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

// DELETE /address
type DeleteAddressReq struct {
	SolAddress string `json:"sol_address"`
}
