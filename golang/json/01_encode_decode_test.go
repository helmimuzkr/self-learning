package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSimplyEncode(t *testing.T) {
	// Contoh data yang akan di Marshal/Encode ke JSON
	nameStr := "Helmi"

	// Untuk encode JSON data menggunakan Marshal function.
	// Jika semuanya berjalan dengan baik, err akan bernilai nil dan data yang sudah di encode akan menghasilkan data mentah yaitu []byte
	// buf(buffer) adalah sebuah variabel sementara yang menampung data []byte JSON
	buf, err := json.Marshal(nameStr)
	if err != nil {
		panic(err)
	}

	// note: buffer adalah tempat menyimpan/memory sementara pada saat perpindahan data

	// buf ini masih raw atau masih dalam bentuk byte.
	// Jika ingin melihat secara jelas harus dikonversi menjadi string terlebih dahulu
	jsonData := string(buf)
	fmt.Println(jsonData)
}

func TestSimplyDecode(t *testing.T) {
	jsonData := []byte(`"Helmi"`)

	// Untuk decode data JSON menggunakan Unmarshal function

	// Buat variable yang nantinya digunakan untuk menampung data yang sudah diUnmarshal
	var name string

	// Ingat, data yang dibuat harus bersifat pointer

	// Lalu, panggil json.Unmarshal.
	// Arg pertama adalah data JSON yang bertipe []byte. ika masih dalam bentuk string, harus dikonversi terlebih dahulu.
	// Arg kedua adalah variabel yang akan dijadikan tempat penyimpanan data yang sudah diUnmarshal.
	// Pada argumen ke dua harus berupa pointer! agar data disimpan ke variabel yang sudah ditunjuk/reference.
	err := json.Unmarshal(jsonData, &name)
	if err != nil {
		panic(err)
	}

	// Jika jsonData menyimpan valid JSON yang cocok dengan name,
	// maka setelah pemanggilan err akan bernilai nil dan data dari jsonData akan tersimpan di dalam variabel name.

	fmt.Println(name)
}
