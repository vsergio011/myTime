package database

import (
	"database/sql" // Interactuar con bases de datos
	"fmt"          // Imprimir mensajes y esas cosas
	"os"

	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

func Open() (*sql.DB, error) {
	usuario := "dev-user"
	pass := "dev-password"
	host := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	if os.Getenv("DB_HOST") == "" {
		host = "tcp(db:3306)" /* "tcp(127.0.0.1:3306)"*/
	}
	fmt.Println("conexion host" + host)
	nombreBaseDeDatos := "myTime"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos)+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
