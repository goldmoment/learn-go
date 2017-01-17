package utils

import (
	"path/filepath"

	"github.com/nu7hatch/gouuid"
)

func GetImageName(filename string) string {
	id, err := uuid.NewV4()
	if err != nil {
		return filename
	}

	return id.String() + filepath.Ext(filename)
}
