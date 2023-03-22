package concurrency

import (
	"fmt"
	"testing"
)

func sum(x, y int, c chan int) {
	data := x + y

	// send data value to channel
	c <- data
}

func TestChan(t *testing.T) {
	// membuat channel yang menerima value dengan tipe data int
	var channel = make(chan int)
	// tutup channel ketika data dalam channel selesai diambil.
	// sangat direkomendasikan setelah menggunakan channel langsung ditutup
	defer close(channel)

	go sum(1, 2, channel)

	// recieve value from channel and store to result
	result := <-channel
	fmt.Println(result) // 3

}

// Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima ini adalah goroutine yang berbeda.
// Ketika proses pengiriman atau penerimaan value goroutine akan terkunci
// hingga sisi yang lain siap untuk mengirim atau menerima.
func TestBlocking(t *testing.T) {
	ch := make(chan string)

	// Terlihat disini string "helmi" dikirim ke ch.
	ch <- "helmi"

	// Lalu, tidak ada variable yang menerima value yang ada di dalam ch.
	// Hal ini akan menyebabkan aplikasi terkunci atau stuck.

	fmt.Println("line ini tidak akan muncul ke layar karena deadlock/block")
}

func TestNonBlock(t *testing.T) {
	ch := make(chan string)

	// Anonymous func ini menerima data dari ch.
	go func() {
		// Data dikirimkan ke ch
		ch <- "helmi"
	}()

	// Pada proses penerimaan data dari ch, maka akan terjadi blocking.
	// Karena msgChannel akan menunggu data yang dikirim ke ch.
	msgChannel := <-ch
	fmt.Println(msgChannel)

	fmt.Println("line ini berhasil muncul")
}
