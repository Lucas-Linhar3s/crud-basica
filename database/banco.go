package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //banco de dados
)

func Conectar() (*sql.DB, error) {
	conexao := "lucas:Go7/flo2@/curso_udemy?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", conexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
