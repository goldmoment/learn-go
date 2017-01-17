package parameter

import ()

type (
    Product struct {
		ID         string    `json:"id"`
		Quantity   int       `json:"quantity"`
	}
	Checkout struct {
		Products   []Product `json:"product"`
	}
)