package uji_coba

import (
	"fmt"
	"testing"
)

type Peripheral interface {
	Input(i string)
	Output()
}

// Struct Keyboard
type Keyboard struct {
	Words string
}

func (k *Keyboard) Input(i string) {
	k.Words = i
}
func (k *Keyboard) Output() {
	fmt.Println(k.Words)
}

// Sturct Mouse
type Mouse struct {
	Click string
}

func (m *Mouse) Input(c string) {
	m.Click = c
}
func (m *Mouse) Output() {
	fmt.Println(m.Click)
}

func TestInterfaceAndMethods(t *testing.T) {
		// Membuat insialisasi interface peripheral
		var peripheral Peripheral

		// NB
		// Struct == Objek
		
		// keyboard := Keyboard{} // Akan error, karena method Keyboard menggunakan pointer reciever.
		// Jika salah satu method pada sebuah struct menggunakan pointer reciever, 
		// maka saat instansiasi struct tersebut harus dengan pointer.
		keyboard := new(Keyboard)
	
		// Assign struct keyboard ke interface peripheral.
		// Akan error jika struct Keyboard bukan pointer.
		peripheral = keyboard
		peripheral.Input("hallo")
		peripheral.Output()
		
		mouse := Mouse{}
		mouse.Input("Kanan")
		mouse.Output()
}