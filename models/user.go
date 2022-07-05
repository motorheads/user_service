package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Country    string `json:"country"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}
