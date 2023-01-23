package materi

import (
	"fmt"
)

func array() {

	// deklarasi array wajib diberi panjangnya
	// array[leng]typedata
	var arr [5]int 

	for i := 0; i < 5; i++ {
		// menambahkan value kedalam array
		// array[index] = value
		arr[i] = i + 1
	}

	// mengambil array dengan index 0
	// array[index]
	fmt.Println(arr)
}