package materi

import "fmt"

var Arr = []string{
    "helmi",
    "muzakir",
    "muhammad",
    "zuwy",
} 

func slices() {

    // arr[low:high]
    var slice = Arr[0:2]
    fmt.Println("Array slice : ", slice)
    
    // leng mulai dari index 1 sampai index 2
    fmt.Println("Panjang : ", len(slice))

    // cap mulai dari index 2 sampai 3
    fmt.Println("Kapasitas : ", cap(slice))
    fmt.Println("")

    // menambahkan value baru kedalam arr pada index terakhir slice
    var appendslice = append(slice, "zumy?")
    fmt.Println("Array append slice : ", appendslice)
    // hasil
    fmt.Println("Array update : ", Arr)
    fmt.Println("")

}