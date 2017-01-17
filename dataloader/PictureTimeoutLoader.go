package dbl

import (
	"time"
	// "fmt"

	"github.com/goldmoment/manager"
	"github.com/goldmoment/model"
)

func AddPictureTimeout(pic *model.Picture) error {
	stmt, err := db.Database.Prepare(`INSERT INTO picture_timeout
	                                  (pictureid, deadline) 
	                                  VALUES(?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	deadline := time.Now().Add(1e9 * 60 * 10) // 10 minutes
	_, err = stmt.Exec(pic.ID, deadline)
	if err != nil {
		return err
	}

	return nil
}

func GetPictureTimeouts(deadline time.Time) []model.PictureTimeout {
	results := []model.PictureTimeout{}

	rows, err := db.Database.Query(`SELECT pt.id, p.path, pt.pictureid 
									FROM picture_timeout AS pt
									INNER JOIN pictures AS p ON pt.pictureid = p.id
									WHERE pt.deadline <= ? 
									LIMIT 0 , 30`, time.Now())
	if err != nil {
		return results
	}

	defer rows.Close()

	for rows.Next() {
		var p model.PictureTimeout
		if err := rows.Scan(&p.ID, &p.Path, &p.PictureID); err != nil {
			break
		}
		results = append(results, p)
	}

	// if err := rows.Err(); err != nil {
	//     log.Fatal(err)

	return results
}

func RemovePictureTimeout(pt model.PictureTimeout) error {
	stmt, err := db.Database.Prepare(`DELETE FROM picture_timeout WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute sql query
	if _, err := stmt.Exec(pt.ID); err != nil {
		return err
	}
	return nil
}

func RemovePictureTimeoutByPicID(pictureId string) error {
	stmt, err := db.Database.Prepare(`DELETE FROM picture_timeout WHERE pictureid = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute sql query
	if _, err := stmt.Exec(pictureId); err != nil {
		return err
	}
	return nil
}
