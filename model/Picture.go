package model

import (
	"time"
)

type (
	Picture struct {
		ID          string    `json:"id"`
		Path        string    `json:"path"`
		Description string    `json:"description"`
		Color       string    `json:"color"`
		Width       int       `json:"width"`
		Height      int       `json:"height"`
		Ratio       float32   `json:"ratio"`
		CreatedAt   time.Time `json:"created_at"`
		ModifiedAt  time.Time `json:"modified_at"`
	}
)
