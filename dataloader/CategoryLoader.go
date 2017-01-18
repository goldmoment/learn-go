package dbl

import (
	// "log"

	"github.com/goldmoment/learn-go/manager"
	"github.com/goldmoment/learn-go/model"
)

func GetCategories(userid string) []model.Category {
	results := []model.Category{}

	rows, err := db.Database.Query(`SELECT c.id, c.title
                    FROM categories c
                    INNER JOIN user_category uc ON c.id = uc.categoryid
                    AND uc.userid = ?
                    LIMIT 0 , 10`, userid) // Start from 0, limit 10 result
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			break
		}
		results = append(results, c)
	}

	return results
}

func GetPublicCategories() []model.Category {
	results := []model.Category{}

	rows, err := db.Database.Query(`SELECT c.id, c.title
                    FROM categories c
                    LIMIT 0 , 10`) // Start from 0, limit 10 result
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			break
		}
		results = append(results, c)
	}

	return results
}
