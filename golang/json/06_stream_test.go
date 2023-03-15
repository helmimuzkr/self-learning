package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// ● Sebelumnya kita belajar package json dengan melakukan konversi data JSON yang sudah dalam bentuk variable dan data string atau []byte
// ● Pada kenyataanya, kadang data JSON nya berasal dari Input berupa io.Reader (File, Network, Request Body)
// ● Kita bisa saja membaca semua datanya terlebih dahulu, lalu simpan di variable, baru lakukan konversi dari JSON,
// 	 namun hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur untuk membaca data dari Stream

// ● Untuk membuat json Decoder, kita bisa menggunakan function json.NewDecoder(reader)
// ● Selanjutnya untuk membaca isi input reader dan konversikan secara langsung ke data di Go-Lang, cukup gunakan function Decode(interface{})

type Product struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
}

type Products []Product

func TestStreamEncoder(t *testing.T) {
	// Create file for store JSON data
	file, err := os.Create("product.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create sample data wich will be encoded to JSON
	products := Products{
		{"P1", "Keyboard", 500000},
		{"P1", "Keyboard", 500000},
	}

	// Encoding Process
	encoder := json.NewEncoder(file)
	err = encoder.Encode(products)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encode stream succeed")
}

func TestStreamDecoder(t *testing.T) {
	// Buka file json
	file, _ := os.Open("product.json")
	defer file.Close()

	// Buat variable yang menampung data hasil decode
	products := new(Products)

	// Proses decode
	// Making object decoder and pass an argument that implement io.Reader
	decoder := json.NewDecoder(file)
	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
	// in this argument the product meant v
	err := decoder.Decode(products) // same as Unmarshal
	if err != nil {
		panic(err)
	}

	fmt.Println(products) // &[{P001 Keyboard 500000} {P002 Mouse 350000}]
}
