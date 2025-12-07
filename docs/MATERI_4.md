# Video 4 | Decode JSON

## Decode JSON
- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON
Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan function json.Unmarshal(byte[], interface{})
- Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer

## Kode Program
```go
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
```