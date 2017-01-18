package dbl

import (
	// 	"fmt"
	// 	"time"

	"../manager"
)

func AddProductPicture(productId string, pictureId string, pictureType int) error {
	stmt, err := db.Database.Prepare(`INSERT INTO product_pictures
	                                  (pictureid, productid, type) 
	                                  VALUES(?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pictureId, productId, pictureType)
	if err != nil {
		return err
	}

	return nil
}

// func RemovePicture(pt model.PictureTimeout) error {
// 	stmt, err := db.Database.Prepare(`DELETE FROM pictures WHERE id = ?`)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Execute sql query
// 	if _, err := stmt.Exec(pt.PictureID); err != nil {
// 		return err
// 	}
// 	return nil
// }
