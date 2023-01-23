package main

/* 	Catatan

Defaultnya, golang mengeksekusi tugas secara synchronous.

Untuk membuat sebuah function dijalankan dengan goroutine ditambahkan keyword "go" sebelum function dijalankan
ketika dijalankan akan dilakukan secara concurrent/async,
sehingga menyebabkan hasil tidak berurutan. Keuntungannya program berjalan lebih cepat, karena menggunakan mengunakan beberapa core cpu.

Goroutine tidak bisa mengembalikan value, dibutuhkan channel
*/

import (
	"fmt"
	"testing"
	"time"
)

func sayHello(arg int) {
	fmt.Printf("perulangan ke-%d\n", arg)
}

func TestTanpaGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		sayHello(i)
	}
}

func TestWithGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go sayHello(i)
	}

	time.Sleep(5 * time.Second)
}