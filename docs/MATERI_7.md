# Video 7 | Map

## Map
- Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
- Artinya atribut nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap
- Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct, kita harus menentukan semua atribut nya
- Untuk kasus seperti ini, kita bisa menggunakan tipe data map[string]interface{}
- Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
- Namun karena value berupa interface{}, maka kita harus lakukan konversi secara manual jika ingin mengambil value nya
- Dan tipe data Map tidak mendukung JSON Tag lagi

## Kode Program
```go
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
```