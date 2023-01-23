package main

import (
	"fmt"
	"testing"
)

func sum(x, y int, c chan int) {
	data := x + y

	// send data to channel
	c <- data
}

func TestChan(t *testing.T) {
	// membuat channel yang menerima value dengan tipe data int
	var channel = make(chan int)
	// tutup channel ketika data dalam channel selesai diambil.
	// sangat direkomendasikan setelah menggunakan channel langsung ditutup
	defer close(channel)

	go sum(1, 2, channel)

	// send data in channel to result
	result := <- channel 
	fmt.Println(result) // 3
	
}