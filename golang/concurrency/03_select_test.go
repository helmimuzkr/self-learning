package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// Deadlock = karena tidak ada yang menerima atau menunggu data channel

// SELECT akan membuat sebuaah goroutine menunggu untuk beberapa operasi komunikasi.
func TestSelect(t *testing.T) {
	ch := make(chan string)

	// Bisa dilihat pada anonymous function dibawah ini.
	// Jika num adalah genap, maka kirim "genap" ke ch.
	// Tetapi num adalah bilangan ganjil, karena valuenya 3.
	// Hal ini akan membuat deadlock/stuck, karena proses pengiriman value ke ch hanya terjadi
	// ketika bilngan num adalah genap.
	go func() {
		num := 3
		if num%2 == 0 {
			ch <- "genap"
		}
	}()

	// Dengan menggunakan primitive select, kita bisa membuat beberapa kondisi
	// apakah ada data yang dikirim ke channel ch atau tidak.
	select {
	// Jika ada yang mengirim data ke channel ch, maka akan diterima oleh variable result dan mentrigger kondisi ini
	// lalu menampilkan result.
	case result := <-ch:
		fmt.Println(result)

	// Jika tidak ada data yang dikirim ke channel ch,
	// yang seharusnya terjadi deadlock karena tidak ada yang nerima data di ch
	// ini malah akan mengeksekusi println.
	// Sebenernya ini lebih ke menunggu. kalau menunggu lebih dari 1 second maka akan mengeksekusi println.
	case <-time.After(1 * time.Second):
		fmt.Println("tidak ada data yang diterima pada channel")
	}

	// Deadlock akan terjadi ketika menunggu terus atau stuck.
}
