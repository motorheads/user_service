package storage

import (
	"fmt"

	"github.com/motorheads/user_service/config"
	"github.com/motorheads/user_service/models"
)

func GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	query := `
			SELECT * 
			FROM users
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Country,
			&user.Address,
			&user.PostalCode,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil

}

func GetUser(user_id int) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf(`
			SELECT *
			FROM users
			WHERE id=%d`, user_id)
	row := config.DB.QueryRow(query)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Country,
		&user.Address,
		&user.PostalCode,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil

}

func CreateUser(user *models.User) error {
	query := `
			INSERT INTO users(
				name,
				email,
				phone,
				country,
				address,
				postal_code
			) VALUES(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
			);
	`
	_, err := config.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Phone,
		user.Country,
		user.Address,
		user.PostalCode,
	)
	return err
}

func UpdateUser(user models.User) error {
	query := `
			UPDATE users 
			SET name=$1,
				phone=$2,
				country=$3,
				address=$4,
				postal_code=$5
			WHERE id=$6
	`

	_, err := config.DB.Exec(
		query,
		user.Name,
		user.Phone,
		user.Country,
		user.Address,
		user.PostalCode,
		user.ID,
	)
	return err
}

func DeleteUser(user_id int) error {
	query := `
			DELETE FROM users
			WHERE id=$1
	`
	_, err := config.DB.Exec(query, user_id)
	return err
}
