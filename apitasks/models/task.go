package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/vsergio011/apitasks/database"
)

type Task struct {
	Id       int64
	Title    string
	Date     string
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
	sql := `SELECT id,title,date,uid_user FROM task WHERE uid_user = ?`

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

func InsertTask(task *Task) (*Task, error) {
	fmt.Println("Estoy en el model")
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlInsertUser := `INSERT INTO task (title,date,uid_user,place) VALUES ( ?, ?, ?,?);`
	fmt.Println(sqlInsertUser)
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	fmt.Println(time.Now().UTC())
	data, err := stmt.Exec(task.Title, task.Date, task.Uid_user, "1")
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	task.Id, err = data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return nil, err
	}

	return task, nil
}

func UpdateTask(task *Task) (*Task, error) {
	fmt.Println("Estoy en el model")
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlInsertUser := `UPDATE task SET title =?, date=?  WHERE id =?;`
	fmt.Println(sqlInsertUser)
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	fmt.Println(time.Now().UTC())
	data, err := stmt.Exec(task.Title, task.Date, task.Id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	task.Id, err = data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return nil, err
	}

	return task, nil
}

func DeleteTask(task *Task) (*Task, error) {
	fmt.Println("Estoy en el model delete")
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlInsertUser := `DELETE FROM task WHERE id = ?;`
	fmt.Println(sqlInsertUser)
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	fmt.Println(task.Id)
	data, err := stmt.Exec(task.Id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return nil, err
	}
	task.Id, err = data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return nil, err
	}

	return task, nil
}
