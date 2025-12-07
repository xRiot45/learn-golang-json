package learngolangjson

import (
	"encoding/json"
	"testing"
)

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
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
