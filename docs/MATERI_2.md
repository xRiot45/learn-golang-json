# Video 2 | Encode JSON

## Encode JSON
- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function json.Marshal(interface{})
- Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal

## Kode Program
```go
package learngolangjson

import (
	"encoding/json"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Thomas")
	logJson(1)
	logJson(true)
	logJson([]string{"a", "b", "c"})
	logJson(map[string]interface{}{"name": "Thomas", "age": 20})
}
```