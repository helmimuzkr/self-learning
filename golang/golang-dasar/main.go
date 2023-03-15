package main

import "fmt"

type User struct {
	Name string
}

type DataUser struct {
	Users []User
}

func (d *DataUser) GetData(u User) {
	
	fmt.Println("didalem func", &u)
	d.Users = append(d.Users, u)

/* 	for _, v := range u {
		d.Users = append(d.Users, v)
	}  */
}

func main() {
	user1 := User{"helmi"}
	/* user2 := User{"muzakir"} */

	data := DataUser{}
	fmt.Println( data)
	data.GetData(user1)
	fmt.Println(data)
}
/* 	data := DataUser{}
	fmt.Printf("Atas : %p\n", user1)
	data.GetData(user1)
	fmt.Printf("Bawah : %p\n", user1)
} */