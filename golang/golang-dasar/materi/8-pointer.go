package materi

import "fmt"

/* Catatan

& (ampersand) = point reference atau mengambil alamat atau pointer
* (asterik) = dereference pointer. jadi value

1. Pointer menyimpan alamat memori dari suatu nilai, bukan nilai itu sendiri.
Untuk mendeklarasikan variabel pointer, Anda harus menggunakan dereferencing operator (*) dan address operator (&).
Contoh:
	var x type = value	=> deklarasi reference
	var y *type = &x	=> deklarasi pointer yang menunjuk ke x
or
	y2 := &x

2. Di Go, kalian bisa menggunakan shorthand (:=) untuk mendeklarasikan variabel pointer dan kalian bisa menggunakan
slice of array atau new() pada struct untuk mengembalikan pointer. Pada struct cukup ditambahkan asterik pada typenya

Contoh :
	var y3 *StructName = new(StructName)
*/

type Mhs struct {
	Nama string
	Npm int
}

// QUIZ
type Gamer struct {
	Name string
	Games []string
}

// code here
func (gamer *Gamer) add(game string) {
	gamer.Games = append(gamer.Games, game) // append/menambahkan data baru(game) kedalam slice gamer.Games
}

func pointers() {

	var x int = 10
	fmt.Println("Value x = ", x)
	var y = &x // menunjuk ke address memory x
	*y = 10 // set value x lewat pointer y
	fmt.Println("value y", *y) // lihat value terbaru melalui pinter y

	// fmt.Println("menampilkan memori variabel 'x' = ", &x)

	
	fmt.Println("")
	

	// jika seperti ini value referensi tidak akan berubah, yang ada malah duplikat
	helmi := 22
	fmt.Println(helmi)

	muzakir := helmi
	muzakir = 11
	// saat dirubah malah hanya variabel muzakir saja yang berubah
	fmt.Println(muzakir)

	
	fmt.Println("")
	

	mhs := Mhs {Nama: "helmi", Npm: 54418672, }	// reference
	fmt.Println("Variable mhs = ", mhs)

	var maba *Mhs	// data baru
	fmt.Println("Variabel Pointer maba = ", maba)

	// bisa juga seperti ini
	kating := &Mhs{"helmi", 1}	// set value mhs dari pointer kating 
	fmt.Println("Variabel Pointer kating = ", kating, mhs)


	
	// Quiz, Buat method dengan menyimpan atau menambahkan game baru kedalam field Games
	gamer := Gamer{Name: "helmi"}
	gamer.add("Apex")
	gamer.add("pubg")
	
	fmt.Println("Daftar Game : ")
	for _, g := range gamer.Games {
		fmt.Println("\t-",g)
	}

	test := &gamer
	test.Name = "muzakir"
	fmt.Println(gamer)
	fmt.Println(test)

	var slices []interface{}
	change := append(slices, "helmi", 1, 2)
	fmt.Println(change)


/* 	error
	*maba := &mhs
	fmt.Println(mhs, maba) */

/*	error 
	*maba := &mhs
	fmt.Println(mhs, *maba) */

/* 	bisa 
	maba := &mhs
	fmt.Println(mhs, *maba) */

}