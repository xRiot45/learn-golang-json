package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Laptop","price":1000000}`
	jsonBytes := []byte(jsonString)

	result := &map[string]interface{}{}
	err := json.Unmarshal(jsonBytes, result)
	if err != nil {
		panic(err)
	}

	fmt.Println((*result)["id"])
	fmt.Println((*result)["name"])
	fmt.Println((*result)["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{"id": 1, "name": "Laptop", "price": 1000000}
	bytes, _ := json.Marshal(product)
	println(string(bytes))
}
