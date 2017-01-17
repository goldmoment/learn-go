package model

import ()

type (
	Product struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
        Price    int    `json:"price"`
        Quantity int    `json:"quantity"`
	}
)