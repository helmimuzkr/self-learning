package oop

import (
	"fmt"
	"testing"
)

// class product
type Product struct {
	// property, representasi data dari sebuah objek
	Judul string 
	Author string
	Publisher string
	Price int32

}

// method dari class product
// method, representasi prilaku/behavior dari sebuah objek




func TestOop(t *testing.T) {
	// instance dari class, jadinya sebuah object baru
	// object := className{Property: value}
	product1 := Product{
		Judul: "Naruto",
		Author: "Masashi Kisimoto",
		Publisher: "Shounen Jump",
		Price: 10000,
	}
	product2 := Product{
		Judul: "God Of War",
		Author: "Eric Williams",
		Publisher: "Sony Interactive Entertainment",
		Price: 500000,
	}
	
	fmt.Println("Product 1: ", product1)
	fmt.Println("Judul Product2: ". product2.Judul)
}

