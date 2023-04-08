package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func readMsg(msgCh chan string) {
	// 3. Lakukan for loop supaya program tidak langsung berhenti.
readLoop: // you can make custom "for" name
	for {
		// 4. Gunakan select supaya program bisa menunggu data yang dikirim ke channel "msgCh".
		select {
		// Ketika ada data yang dikirim ke "msgCh"
		// maka akan membaca message tersebut dan hentikan program "readMsg".
		case msg := <-msgCh:
			fmt.Println("read message")
			fmt.Println(msg)
			fmt.Println("read message done.")
			break readLoop // break the "readLoop"

		// Jika tidak ada data yang dikirim ke channel "msgCh",
		// Freeze 1 detik dan print "waiting message...".
		// Maka, selama 5 detik kondisi default ini akan ditrigger terus menerus karena for loop.
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("waiting message...")
		}
	}
}

func TestForSelect(t *testing.T) {
	// 1. Membuat channel message, yang nantinya digunakan untuk komunikasi antar routine.
	msgCh := make(chan string)
	defer close(msgCh) // close the channel after function done.

	// 2. Eksekusi function "readMsg" dengan goroutine.
	go readMsg(msgCh)

	// 5. Buat time sleep(freeze) selama 5 detik supaya for select mentrigger kondisi default.
	// Jadi, seperti menunggu pesan selama 5 detik.
	time.Sleep(5 * time.Second)

	// 6. Setelah menunggu selama 5 detik, masukan data "halo helmi!" ke dalam channel.
	msgCh <- "halo helmi!"
}
