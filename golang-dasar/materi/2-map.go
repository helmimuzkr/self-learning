package materi

import (
	"fmt"
)

/* Catatan

map sama seperti objek pada di bahasa pemrograman lain.

	1. membuat map => var variable map[typeKey]typeValue
	2. membuat map dengan value campuran => var variable map[typeKey]interface{}
	3. menambahkan data dalam map => map[key] = value
	4. delete data dalam map => delete(map, key)

*/

func maps() {
	// map dengan value campuran
	human := map[string]interface{} {
		"nama": "helmi",
		"umur": 22,
		"isKerja": false,
	}

	// menambahkan data baru ke dalam map human
	human["alamat"] = "bekasi"

	// hapus data
	// delete(human, ["alamat"])

	fmt.Println(human)
}