package materi

import "fmt"

/* Catatan
Struct adalah tipe data abstract yang berisi kumpulan variabel atau field.

1. pembuatan type Struct
2. format penamaan struct dan field menggunakan CapitalCase
3. format pembuatan type struct
	type NamaStruct struct {Field type}
4. format pembuatan struct
	name := NamaStruct{Field: value}
5. Format pembuatan struct method
	func (name NamaStruct ) nama_function(){}
*/

// struct or object
type User struct {
	Id int
	Name string
	Email string
	IsActive bool
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

// methods
// dimiliki atau di dalam struct User
func (u User) Introduction() {
	fmt.Println("Hallo nama saya", u.Name)
}
/* function dengan parameter struct 
	Bedanya function itutidak dimiliki oleh siapa pun
	
	// pembuatan
	func intro(u User) {
		fmt.Println("Hallo nama saya", u.Name)
	} 

	// akses
	intro(user1)
*/

func (g Group) GetInfoGroup() {
	fmt.Println("Nama group :", g.Name)
	fmt.Println("Admin : ", g.Admin.Name)
	fmt.Println("Anggota : ")

	for _, user := range g.Users {
		fmt.Println("\t", user.Name)
	}
}

func structs() {
	// structs or objects
	// cara 1
	var user1 User
	user1 = User{1, "helmi", "h@gmail.com", true}
	// cara 2
	var user2 User = User{2, "muzakir", "m@gmail.com", true}
	// cara 3, Best Practice
	user3 := User{3, "muhammad", "muh@gmail.com", false}

	users := []User{user1, user2, user3}

	group1 := Group{"e-sport", user1, users, true}

	fmt.Println(group1)

	// methods
	user1.Introduction() 

	// Soal. Buat method untuk group
	group1.GetInfoGroup()

}