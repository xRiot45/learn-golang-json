# Video 5 | JSON Array

## JSON Array
- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
- Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks (Object atau Array)
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice

## Decode JSON Array
-   Selain menggunakan Array pada attribute di Object
- Kita juga bisa melakukan encode atau decode langsung JSON Array nya
- Encode dan Decode JSON Array bisa menggunakan tipe data Slice

## Kode Program
```go
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
```