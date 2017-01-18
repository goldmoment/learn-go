package dbl

import (
	// "fmt"
	"time"

	"github.com/goldmoment/learn-go/manager"
	"github.com/goldmoment/learn-go/model"

	"github.com/nu7hatch/gouuid"
)

func AddPicture(pic *model.Picture) error {
	stmt, err := db.Database.Prepare(`INSERT INTO pictures
	                                  (id, path, description, color, width,
	                                   height, ratio, created_at, modified_at) 
	                                  VALUES(?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	pic.ID = id.String()
	pic.CreatedAt = time.Now()
	pic.ModifiedAt = pic.CreatedAt

	_, err = stmt.Exec(pic.ID, pic.Path, pic.Description, pic.Color,
		pic.Width, pic.Height, pic.Ratio, pic.CreatedAt, pic.ModifiedAt)
	if err != nil {
		return err
	}

	return nil
}

func RemovePicture(pt model.PictureTimeout) error {
	stmt, err := db.Database.Prepare(`DELETE FROM pictures WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute sql query
	if _, err := stmt.Exec(pt.PictureID); err != nil {
		return err
	}
	return nil
}

func UpdatePicture(pictureID string, description string) error {
	stmt, err := db.Database.Prepare(`UPDATE pictures SET description = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute sql query
	if _, err := stmt.Exec(description, pictureID); err != nil {
		return err
	}

	return nil
}
