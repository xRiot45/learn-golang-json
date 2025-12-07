package learngolangjson

import (
	"encoding/json"
	"testing"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func TestJSONTag(t *testing.T) {
	product := Product{1, "Laptop", 1000000}
	bytes, _ := json.Marshal(product)
	println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Laptop","price":1000000}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	println(product.Id)
	println(product.Name)
	println(product.Price)
}
