package models

import (
	"errors"
	"time"

	"github.com/vsergio011/apitasks/database"
)

type Task struct {
	Id       int64
	Title    string
	Date     time.Time
	Uid_user string
}

func GetTasks() ([]Task, error) {
	sql := `SELECT id,title,date,uid_user FROM task`

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
		if err := rows.Scan(&b.Id, &b.Title, &b.Date, &b.Uid_user); err != nil {
			continue
		}
		tasks = append(tasks, b)
	}

	return tasks, nil
}

func GetTask(id int64) (*Task, error) {
	sqlGetUserCredentials := `SELECT id,title,date,uid_user FROM task WHERE id=?`
	db, err := database.Open()
	if err != nil {
		return nil, errors.New("error opening the database")
	}
	defer db.Close()

	var taskDB Task
	err = db.QueryRow(sqlGetUserCredentials, id).Scan(&taskDB.Id, &taskDB.Title, &taskDB.Date, &taskDB.Uid_user)
	if err != nil {
		return nil, errors.New("error executing query")
	}

	return &taskDB, nil
}
func GetTaskscreatedBy(uid_user string) ([]Task, error) {
	sql := `SELECT id,title,date,uid_user FROM task WHERE id_user = ?`

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasks []Task
	rows, err := db.Query(sql, uid_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Task
		if err := rows.Scan(&b.Id, &b.Title, &b.Date, &b.Uid_user); err != nil {
			continue
		}
		tasks = append(tasks, b)
	}

	return tasks, nil
}

//Sin probar
func insertTask(task *Task) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlInsertUser := `INSERT INTO task (id,title,date,id_user) VALUES (?, ?, ?, ?);`
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	data, err := stmt.Exec(task.Id, task.Title, task.Date, task.Uid_user)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}
	task.Id, err = data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return err
	}

	/*sqlInsertUserCredentials := `INSERT INTO users_credentials (user, password, id_user) VALUES (?, ?, ?);`
	stmt, err = db.Prepare(sqlInsertUserCredentials)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT INSERT USER CREDENTIALS")
		return err
	}

	data, err = stmt.Exec(signInUser.Username, signInUser.Password, signInUser.Id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT INSERT USER CREDENTIALS")
		return err
	}*/

	return nil
}
