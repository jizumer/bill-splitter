package main

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (p *Person) Create() error {
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(name, balance) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(p.Name, p.Balance)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}

	return nil
}

func (p *Person) FindAll() ([]Person, error) {
	db := GetConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	person := []Person{}
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.Name, &p.Balance)
		if err != nil {
			return nil, err
		}
		person = append(person, p)
	}
	return person, nil
}

func (p *Person) Update() error {
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE person SET name=?, balance=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(p.Name, p.Balance, p.ID)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}

	return nil
}
