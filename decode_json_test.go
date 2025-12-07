package learngolangjson

import (
	"encoding/json"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Firstname":"Thomas","Middlename":"Wijaya","Lastname":"Agung","Age":20}`
	jsonBytes := []byte(string(jsonString))

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	println(customer.Firstname)
	println(customer.Middlename)
	println(customer.Lastname)
	println(customer.Age)
}
