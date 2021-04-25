package main

import (
	"database/sql"
	"errors"
)

var ErrNoRow = errors.New("no rows")

type MyError struct{
	Error error
	Sql string
}


func error_wrap(db *sql.DB) (string, error){
	sql := "select name from user where id = ? "
	rows, err := db.Query(sql, 1)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		return nil, &MyError{
			Error : err,
			Sql : sql
		}
	}

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return name, nil
}
