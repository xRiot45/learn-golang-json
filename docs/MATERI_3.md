# Video 3  | JSON object

## JSON Object
- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
- Sedangkan value nya baru berupa 

## Struct
- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct

## Kode Program
```go
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
```