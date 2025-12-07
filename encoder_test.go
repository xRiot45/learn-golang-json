package learngolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Firstname: "Thomas",
		Lastname:  "Alberto",
	}

	encoder.Encode(customer)

}
