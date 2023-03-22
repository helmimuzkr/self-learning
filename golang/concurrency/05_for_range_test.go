package concurrency

import (
	"fmt"
	"testing"
)

// for - range ada

func sendMessage(ch chan<- string) {
	for i := 0; i < 5; i++ {
		fmt.Println("send message...", i+1)
		msg := fmt.Sprintf("Hallo ini pesan ke- %d", i)
		ch <- msg
	}
	// Jika tidak ditutup akan deadlock
	fmt.Println("tutup channel")
	close(ch)
}

func printMessage(ch <-chan string) {
	for msg := range ch {
		fmt.Println("read message...")
		fmt.Println(msg)
	}
}

func TestForRange(t *testing.T) {
	ch := make(chan string)

	fmt.Println("function send message dipanggil")
	go sendMessage(ch)

	fmt.Println("function read message dipanggil")
	printMessage(ch)
}
