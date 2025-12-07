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
