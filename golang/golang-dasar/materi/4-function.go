package materi

import (
	"errors"
	"fmt"
	"math"
)

// var name = []string {"helmi", "muzakir", "muhammad"}

// function with params
// func name(params typedata) typedata{}
func sum(a, b int) int{
	return a + b
}

// function mengembalikan 2 string
func swap(a, b string) (string, string) {
	return a, b
}

// variadic func
// func name(params ...type) none or type {}
func variadic(names ...string) {
	for _, name := range names {
		fmt.Printf("%s ", name)
	}
	fmt.Println("")
}

// function as parameter
func nameFilter(name string, filter func(string) string) {
	fmt.Println("Hallo ", filter(name))
}
func filter(name string) string {
	if name == "anjing" {return "..."}
	return name
}

// function as a type
// type nameType func(typeParams)typeReturn
type Filter func(int) string

func oddEvenCheck(num int, filter Filter) {
	fmt.Println(filter(num))
}

func check(num int) string {
	if num % 2 == 0 {
		return "Bilangan Genap"
	}
	return "Bilangan ganjil"
}

// TEST
// penjumlahan dengan input array
func Sum(numbs[]int) int {
	var result int 

	for _, value := range numbs {
		result += value
	}

	return result
}
/* // Bisa menggunakan variadics func
func Sum(numb ...int) int {
	var result int
	for _, value := range numb {
		result = result + value
	}

	return result
}
*/

// menghitung dengan inputan operasi
func Calc(a int, b int, opt string) (int, error) {
	var result int
	var err error

	switch opt {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		err = errors.New("Operator tidak dikenali")
	}

	return result, err
}

func function() {
	fmt.Println(math.Pi)

	fmt.Println(sum(1, 2))

	fmt.Println(swap("hallo", "helmi"))

	// default calling variadic func
	variadic("muhammad", "helmi", "muzakir")
	// array as argument
	var name = []string {"helmi", "muzakir", "muhammad"}
	variadic(name...)

	// function as params
	nameFilter("helmi", filter)

	// function as type
	oddEvenCheck(1, check)


	// TEST
	// soal 1
	numbs := []int {10, 5, 3}
	fmt.Println(Sum(numbs))

/* 	// Bisa menggunakan variadics func
	numb := []int {10, 5, 3}
	fmt.Println(sum(numb...))
	*/

	// soal 2
	result, err := Calc(10, 2, "+") // 12

	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(result)
	}
	
}