package model

import ()

type (
	PictureTimeout struct {
		ID        int    `json:"id"`
		Path      string `json:"path"`
		PictureID string `json:"pictureid"`
	}
)
