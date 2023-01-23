package golang_httprouter

import (
	"fmt"
	"testing"
)

type People interface {
	sayHello(string)
}

type greetFunc func(string)

func (g greetFunc) sayHello(n string) {
	fmt.Println("sayHello")
	g(n)
}

func filter(p People) People {
	fmt.Println("Function filter dijalankan")

	// function belum dijalankan, hanya dijadikan return value
	return greetFunc(func(n string) {
		fmt.Println("filtering 1...")
		p.sayHello(n)
	})
}

func filter2(p People) People {
	fmt.Println("Function filter 2 dijalankan")

	// function belum dijalankan, hanya dijadikan return value
	return greetFunc(func(n string) {
		fmt.Println("filtering 2...")
		p.sayHello(n)
	})
}

func TestCoba(t *testing.T) {
	salam := greetFunc(func(n string) {
		fmt.Println("Hallo", n)
	})


	var people People = salam // people(n string) { fmt.Println("hallo", n) }
	people = filter(people) // people(n string) { p.sayHello(n) }

	fmt.Println("assign filter2(people) ke people")

	people = filter2(people) // people(n string) { p.sayHello(n) }
	people.sayHello("Helmi") // panggil dan jalankan function

	fmt.Println("")


/* 	test := func(n string) func() string {
		fmt.Println("function test dijalankan", n)
		return func() string {
			//code
			return n + " helmi"
		}
	}
	test2 := test("string")
	fmt.Println(test2()) */
}

// func filter2(p People) People {
// 	fmt.Println("Function filter 2 dijalankan")

// 	// function belum dijalankan, hanya dijadikan return value
// 	return greetFunc(func(n string) {
// 		fmt.Println("filtering 2...")
// 		p.sayHello(n)
// 	})
// }

// type greetFunc func(string)

// func (g greetFunc) sayHello(n string) {
// 	fmt.Println("sayHello")
// 	g(n) == jalanin function ini => greetFunc(func(n string) {
// 											fmt.Println("filtering...")
// 											p.sayHello(n)
// 										})
// }

// greetFunc(func(n string) {
// 	fmt.Println("filtering 1...")
// 	p.sayHello(n)
// })



