package models
import (
	"github.com/vsergio011/apitasks/database"
	//"time"
)
type Task struct {
	Id  int 
	Title string
	Date string
	Id_user int
}

func GetTasks() ([]Task, error) {
	sql := `SELECT id,title,date,id_user FROM task`

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasks []Task
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Task
		if err := rows.Scan(&b.Id, &b.Title, &b.Date, &b.Id_user); err != nil {
			continue
		}
		tasks = append(tasks, b)
	}

	return tasks, nil
}
