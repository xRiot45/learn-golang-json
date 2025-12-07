package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		Firstname:  "Thomas",
		Middlename: "Wijaya",
		Lastname:   "Agung",
		Age:        20,
		Hobbies:    []string{"Coding", "Reading"},
	}

	bytes, _ := json.Marshal(customer)
	println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Firstname":"Thomas","Middlename":"Wijaya","Lastname":"Agung","Age":20,"Hobbies":["Coding","Reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname:  "Thomas",
		Middlename: "Wijaya",
		Lastname:   "Agung",
		Age:        20,
		Hobbies:    []string{"Coding", "Reading"},
		Address: []Address{
			{
				Street:     "Jl. Raya",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "Jl. Raya",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"Thomas","Middlename":"Wijaya","Lastname":"Agung","Age":20,"Hobbies":["Coding","Reading"],"Address":[{"Street":"Jl. Raya","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Raya","Country":"Indonesia","PostalCode":"12345"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Address)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl. Raya","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Raya","Country":"Indonesia","PostalCode":"12345"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
