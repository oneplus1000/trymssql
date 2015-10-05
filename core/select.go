package core

import "fmt"

func SelectFromMsSql() error {

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
