package main

import (
	"database/sql"
	"log"
	"os"
)

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
	log.Println("Database connection established")
	return db
}

func SetupDatabase() {
	if isDatabaseSetup() {
		log.Println("Database is already setup")
		return
	}
	log.Println("Setting up database")
	err := ExecuteSQLScript("init.sql")
	if err != nil {
		log.Fatal(err)
	}

}

func isDatabaseSetup() bool {
	db := GetConnection()

	_, err := db.Query("SELECT * FROM person")
	if err != nil {
		return false
	}
	return true
}

func ExecuteSQLScript(filename string) error {
	db := GetConnection()

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	sql := string(content)
	log.Println(sql)

	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	log.Println("Database setup complete")
	return nil
}
