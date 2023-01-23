package materi

import "fmt"

func fors() {

	var slice = []string {"muhammad", "helmi", "muzakir"}
	var mhs = make(map[string]string)
	
	mhs["nama"] = "helmi"
	mhs["kelas"] = "4IA07"
	mhs["npm"] = "54418672"

	// for biasa
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%s ", slice[i])
		
	}

	fmt.Println("")

	// for range
	for key, value := range mhs{
		fmt.Println(key, "=", value)
	}
	fmt.Println("")
	for index, value := range slice {
		fmt.Printf("index ke = %d, value = %s \n", index, value)
	}
	fmt.Println("")
	/* 
		Jika tidak ingin menampilkan index, maka bisa mengganti parameter pertama dengan "_"
		"_" atau underscore berfungsi untuk memberitahu kalau variabel tidak dibutuhkan atau ignore
		for _, value := range array{}
	*/
	for _, value := range slice {
		fmt.Printf("%s ", value)
	}
	
}