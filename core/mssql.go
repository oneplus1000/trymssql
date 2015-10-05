package core

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func MsSqlOpen() (*sql.DB, error) {
	db, err := sql.Open("mssql", "server=localhost;user id=sa;password=123456;database=godb")

	if err != nil {
		return nil, err
	}
	return db, nil
}
