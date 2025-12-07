# Video 6 | JSON Tag

## JSON Tag
- Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut  yang sama (case sensitive)
- Kadang ada style yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase
- Untungnya, package json mendukun Tag Reflection
- Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan ketika konversi dari atau ke JSON

## Kode Program
```go
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
```