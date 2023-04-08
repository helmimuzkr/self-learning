package concurrency

import (
	"fmt"
	"testing"
)

func TestSelectDefault(t *testing.T) {

	msgCh := make(chan string)

	go func() {
		msgCh <- "halo helmi!"
	}()

	// Jika di Comment akan mentrigger kondisi "default"
	// time.Sleep(10 * time.Millisecond)

	select {
	// Jika ada data yang dikirim ke channel, maka eksekusi kondisi ini
	// dan assign data channel ke variable msg
	case msg := <-msgCh:
		fmt.Println(msg)

	// Jika tidak ditunggu menggunakan time.Sleep dan tidak ada data yang dikirim kee channel,
	// maka akan mentrigger kondisi "default".
	default:
		fmt.Println("belum ada message dikirim")
	}
}
