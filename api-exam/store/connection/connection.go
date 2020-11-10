package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDataBase() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "testDB"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetList(query string, fn func(*sql.Rows) error) error {
	db, err := getDataBase()
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query(query)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		//err = rows.Scan(&c.Id, &c.Nombre, &c.Direccion, &c.CorreoElectronico)
		err = fn(rows)
		if err != nil {
			return err
		}
	}
	// Vacío o no, regresamos el arreglo de contactos
	return nil
}
