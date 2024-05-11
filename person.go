package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (p *Person) FindAll() ([]Person, error) {

	file, err := os.Open("db/persons.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	persons := []Person{}
	err = decoder.Decode(&persons)
	if err != nil {
		return nil, err
	}

	log.Println("Persons: ", persons)
	return persons, nil

}
