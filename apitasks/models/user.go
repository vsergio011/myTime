package models
import (
	"github.com/vsergio011/apitasks/handlers"
)
type User struct {
	Id  int 
	Name string
	Surname string
	Email string
	Token string
}

func GetUsers() ([]User, error) {
	sql := `SELECT users.id, users.name, 
									 users.surname, users.email FROM users`

	database, err := handlers.Open()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var users []User
	rows, err := database.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b User
		//if err := rows.Scan(&b.Id, &b.Model, &b.Address, &b.Lat, &b.Lon, &b.Booked, &b.DateRent, &b.DateReturn); err != nil {
		if err := rows.Scan(&b.Id, &b.Name, &b.Surname, &b.Email); err != nil {
			continue
		}
		users = append(users, b)
	}

	return users, nil
}