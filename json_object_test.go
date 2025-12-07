package learngolangjson

import (
	"encoding/json"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Hobbies    []string
	Address    []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname:  "Thomas",
		Middlename: "Wijaya",
		Lastname:   "Agung",
		Age:        20,
	}
	bytes, _ := json.Marshal(customer)
	println(string(bytes))
}
