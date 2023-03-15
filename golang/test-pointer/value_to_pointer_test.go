package uji_coba

import (
	"fmt"
	"testing"
)

/**
 * Test 1 => Ubah return jadi pointer di construct
 * Test 2 => Ubah reciever pada method GetUser jadi pointer
 * Test 2 => Ubah reciever pada method SetNama jadi pointer
 */

// interface
type UserRepo interface {
	GetUser()
	SetNama(n string)
}

// struct
type User struct {
	Nama string
	Umur int32
}

// construct
func NewUser(n string, u int32) *User {
	// address 
	user := User{
		Nama: n,
		Umur: u,
	}
	// default address 0xc0000081f8
	// setelah test 1, referencenya ngikut ke implementasi objeknya -> 0xc0000081e0
	// setelah test 2, Masih sama 0xc0000081e0
	// setelah test 3, Masih sama 0xc0000081e0
	fmt.Println("di dalam construct : ", &user.Nama)
	fmt.Printf("hallo nama construct %s \n\n", user.Nama)
	return &user
}

// method
func (u *User) GetUser() {
	// address 0xc000008210
	// setelah test 1 -> 0xc0000081f8. Addressnya akan sama seperti memory pada constructor
	// setelah test 2, referencenya ngikut ke implementasi objeknya -> 0xc0000081e0
	// setelah test 3, Masih sama 0xc0000081e0
	fmt.Println("di getUser : ", &u.Nama)
	fmt.Printf("hallo nama getUser %s \n\n", u.Nama)
}

func (u *User) SetNama(s string) {
	u.Nama = s
	// address 0xc000008228
	// setelah test 1 -> 0xc000008210. Addressnya akan sama seperti default address GetUser
	// setelah test 2 -> 0xc0000081f8. Addressnya akan sama seperti address GetUser setelah test 1
	// setelah test 3, referencenya ngikut ke implementasi objeknya -> 0xc0000081e0
	fmt.Printf("\ndi SetNama : %p\n", &u.Nama)
	fmt.Printf("hallo nama setNama %s \n\n", u.Nama)
}

func TestCoba(t *testing.T) {
	user1 := NewUser("Helmi", 23) // reference
	// address 0xc0000081e0
	fmt.Println("address : ", &user1.Nama)
	
	fmt.Println("")

	user1.GetUser()
	// address 0xc0000081e0
	fmt.Println("Sebelum ganti nama : ", user1.Nama)
	fmt.Println("address : ", &user1.Nama)

	user1.SetNama("muzakir")
	// adress 0xc0000081e0
	fmt.Println("setelah ganti nama : ", user1.Nama)
	fmt.Println("address : ", &user1.Nama)
}

