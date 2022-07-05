package storage

import (
	"fmt"

	"github.com/motorheads/catalog_service/config"
	"github.com/motorheads/catalog_service/models"
)

func GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product

	query := `
			SELECT * 
			FROM tires
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Type,
			&product.Price,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil

}

func GetProductByID(product_id int) (*models.Product, error) {
	var product models.Product

	query := fmt.Sprintf(`
			SELECT *
			FROM tires
			WHERE id=%d`, product_id)
	row := config.DB.QueryRow(query)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Type,
		&product.Price,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &product, nil

}

func CreateProduct(product *models.Product) error {
	query := `
			INSERT INTO tires(
				id,
				name,
				type,
				price
			) VALUES(
				$1,
				$2,
				$3,
				$4
			);
	`
	_, err := config.DB.Exec(query, product.ID, product.Name, product.Type, product.Price)
	return err
}

func UpdateProduct(product models.Product) error {
	query := `
			UPDATE tires SET
				id = $1,
				name = $2,
				type = $3,
				price = $4
				where id = $5
	`

	_, err := config.DB.Exec(
		query,
		product.ID,
		product.Name,
		product.Type,
		product.Price,
		product.ID,
	)
	return err
}

func DeleteProductByID(product_id int) error {
	query := `
			DELETE FROM tires
			WHERE id=$1
	`
	_, err := config.DB.Exec(query, product_id)
	return err
}
