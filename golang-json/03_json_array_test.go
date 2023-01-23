package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	// Contoh data array
	product := []string{"Keyboard", "Mouse", "Headset"}

	// Marshal
	buf, errEnc := json.Marshal(product)
	if errEnc != nil {
		panic(errEnc)
	}
	jsonData := string(buf)
	fmt.Printf("Data yang sudah diEncode/Marshal: \n%s\n\n", product) // [Keyboard Mouse Headset]

	// Unmarshal
	var newProduct []string
	errDec := json.Unmarshal([]byte(jsonData), &newProduct)
	if errDec != nil {
		panic(errDec)
	}
	fmt.Printf("Data yang sudah diDecode/Unmarshal: \n%v\n\n", newProduct) // [Keyboard Mouse Headset]
}
