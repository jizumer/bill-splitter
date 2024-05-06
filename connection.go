package main

import "database/sql"

var db *sql.DB

func GetConnection() *sql.DB {

	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}
	return db
}
