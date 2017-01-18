package dbl

import (
	// "log"
	// "fmt"

	"github.com/goldmoment/learn-go/manager"
	"github.com/goldmoment/learn-go/model"

	"github.com/nu7hatch/gouuid"
)

func AddProduct(categoryid string, product *model.Product) error {
	stmt, err := db.Database.Prepare(`INSERT INTO products(categoryid, name, price, quantity, id) VALUES(?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	product.ID = id.String()

	if _, err := stmt.Exec(categoryid, product.Name, product.Price, product.Quantity, product.ID); err != nil {
		return err
	}

	return nil
}

func GetProducts(categoryid string) []model.Product {
	results := []model.Product{}

	rows, err := db.Database.Query(`SELECT p.id, p.name, p.price, p.quantity 
                    FROM products p
                    WHERE p.categoryid = ?
                    LIMIT 0 , 30`, categoryid)
	if err != nil {
		return results
	}

	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
			break
		}
		results = append(results, p)
	}

	// if err := rows.Err(); err != nil {
	//     log.Fatal(err)
	// }

	return results
}

func GetProductsHotest() []model.Product {
	results := []model.Product{}

	rows, err := db.Database.Query(`SELECT p.id, p.name, p.price, p.quantity 
                    FROM products p
                    LIMIT 0 , 4`)
	if err != nil {
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
			break
		}
		results = append(results, p)
	}

	return results
}
