package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ● Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
// ● Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
// ● Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
// ● Sedangkan valuenya berupa interface

type People struct {
	Name string
	Age  int32
}

func TestJsonObject(t *testing.T) {
	// Contoh data yang akan di Marshal/Encode ke JSON
	dataPeople := People{
		Name: "Helmi",
		Age:  23,
	}

	// Marshal/Encode jsonData
	buf, errEnc := json.Marshal(dataPeople)
	if errEnc != nil {
		panic(errEnc)
	}

	// Konversi data json ke string dan simpan ke dalam variable baru
	dataJson := string(buf)
	fmt.Println(dataJson) // {"Name":"Helmi","Age":23}

	// Unmarshal/Decode data JSON
	people := new(People)

	errDec := json.Unmarshal([]byte(dataJson), people)
	if errDec != nil {
		panic(errDec)
	}

	// Jika sudah melakukan proses Unmarshal hasilnya akan langsung berbentuk Struct atau Map,
	// tergantung tipe data penampungnya.
	fmt.Printf("Tipe data: %T\n", people)
	fmt.Printf("Isi data: %+v\n", people) // &{Name:Helmi Age:23}
}
