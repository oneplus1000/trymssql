package core

import (
	"database/sql"
	"fmt"
)

func MsSqlOpen() (*sql.DB, error) {
	db, err := sql.Open("mssql", "server=localhost;user id=sa;password=123456;database=godb")

	if err != nil {
		return nil, err
	}
	return db, nil
}

func ImportFromMsSql() error {

	db, err := MsSqlOpen()
	if err != nil {
		return err
	}
	defer db.Close() //ปิด db

	rows, err := db.Query("select id,name from table01")
	if err != nil {
		return err
	}
	defer rows.Close() //ปิด rows

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("id=%d %s\n", id, name)
	}
	return nil
}
