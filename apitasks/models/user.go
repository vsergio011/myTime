package models
import (
	"github.com/vsergio011/apitasks/database"
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

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b User
		if err := rows.Scan(&b.Id, &b.Name, &b.Surname, &b.Email); err != nil {
			continue
		}
		users = append(users, b)
	}

	return users, nil
}
