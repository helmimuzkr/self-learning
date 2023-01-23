package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ● Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut yang sama (case sensitive)
// ● Kadang ada style yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase
// ● Untungnya, package json mendukun Tag Reflection
// ● Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan ketika konversi dari atau ke JSON

// JSON Tag hanya bisa digunakan pada struct
// JSON Tag best practicenya menggunakan snake_case
// Untuk menggunakannya bisa menambahkan `json:"nama_tag"` disamping fieldnya
type Item struct {
	ItemId string `json:"item_id"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
}

func TestJsonTag(t *testing.T) {
	// Marshal
	item := Item{
		ItemId: "P01",
		Name:   "Keyboard",
		Price:  500000,
	}
	buf, errMarshal := json.Marshal(item)
	if errMarshal != nil {
		panic(errMarshal)
	}
	jsonData := string(buf)
	fmt.Println(jsonData) // {"item_id":"P01","name":"Keyboard","price":500000}

	// Unmarshal
	newItem := new(Item)
	errUnmarshal := json.Unmarshal([]byte(jsonData), newItem)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	fmt.Println(newItem) // &{P01 Keyboard 500000}

	// Pada saat Unmarshal/Read data JSON, golang dapat otomatis menyesuaikan nama atribut JSON ke nama field newItem. Karena tidak case sensitive
	// Kecuali ada simbol pemisah pada nama atribut, maka akan tidak dibaca oleh golang.
	// Misalnya:
	jsonWrong := `{"ItemId":"P01","namE":"Keyboard","PRICE":500000}`
	wrong := new(Item)
	json.Unmarshal([]byte(jsonWrong), wrong)
	fmt.Println(wrong) // &{ Keyboard 500000}

	// namE dan PRICE terbaca, padahal pada tag struct Item lower case semua.
	// Ini dikarenakan golang otomatis menyesuaikan nama atribut JSON dengan nama field struct Item.

	// Berbeda dengan atribut name dan price, atribut ItemId tidak terbaca.
	// Pada json ItemId tidak terbaca dikarenakan pada tag fieldnya ItemId menggunakan underscore(_).

	// Contoh yang benar:
	jsonCorrect := `{"Item_Id":"P01","namE":"Keyboard","PRICE":500000}`
	correct := new(Item)
	json.Unmarshal([]byte(jsonCorrect), correct)
	fmt.Println(correct) // &{P01 Keyboard 500000}
}
