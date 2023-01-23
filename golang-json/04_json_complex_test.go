package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Mahasiswa struct {
	NPM     int32
	Name    string
	Hobbies []string
}

func TestJsonComplex(t *testing.T) {
	// Contoh data
	mahasiswa := &Mahasiswa{
		NPM:  54418672,
		Name: "Helmi",
		Hobbies: []string{
			"Playing Game",
			"Reading",
		},
	}

	// Marshal
	buf, errEnc := json.Marshal(mahasiswa)
	if errEnc != nil {
		panic(errEnc)
	}
	jsonData := string(buf)
	fmt.Printf("After Marshal : \n%s\n", jsonData)

	// Unmarshal
	newMahasiswa := new(Mahasiswa)
	errDec := json.Unmarshal([]byte(jsonData), newMahasiswa)
	if errDec != nil {
		panic(errDec)
	}
	fmt.Printf("After Unmarshal : \n%+v\n", newMahasiswa)
}
