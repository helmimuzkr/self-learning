package materi

import "fmt"

func shotrtStatement() {
	var nama string = "helmi"

	// standard if
	if nama == "helmi" {
		fmt.Println("hallo", nama)
	}else {
		fmt.Println("tidak dikenali")
	}

	// if short statement
	// if statement; condition{}
	if gender := "pria"; gender == "pria" {
		fmt.Println("Dia adalah pria")
	}

}