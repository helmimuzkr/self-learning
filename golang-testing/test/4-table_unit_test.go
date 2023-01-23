package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/* Catatan

Table unit test adalah testing yang biasa para programmer Golang gunakan
untuk melakukan test secara dinamis.

jadi, nanti akan di buat slice of array yang menampung data yang berisi parameter
dan ekspetasi dari hasil test, lalu berikutnya dilakukan eksekusi dengan cara iterasi.cons

untuk membuat sub testharus menggunakan method Run(nama_test, func(t *testing.T){})

*/

// membuat struct untuk test
type Test struct {
	Name string
	Expect int
	Request int
}

func TestTable(t *testing.T) {
	data := &[]Test{
		{
			Name: "Sum(1,2,3)",
			Expect: 6,
			Request: Sum(1,2,3),
		}, 
		{
			Name: "Mul(1,2,3)",
			Expect: 6,
			Request: Mul(1,2,3),
		},
	}
	
	for _, d := range *data {
		t.Run(d.Name, func(t *testing.T) {
			result := d.Request

			require.Equal(t, d.Expect, result)
		})
	}
}