package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang_mysql")
	if err != nil {
		return nil, err
	}

	return db, nil
}
