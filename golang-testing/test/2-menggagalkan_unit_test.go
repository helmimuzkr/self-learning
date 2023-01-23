package helper

/* Catatan

Pada testing, disini ada 4 cara melakukan handle fail test
1. t.Fail()			: tetap melanjutkan eksekusi ketika error
2. t.FailNow()		: menghetikan eksekusi ketika terjadi error
3. t.Error(args)	: seperti Fail(), tetapi memunculkan argumen
4. t.Fatal(args)	: seperti FailNow(), tetapi memunculkan argumen <- direkomendasikan

*/

import (
	"fmt"
	"testing"
)

// t.Fail()
func TestFail(t *testing.T) {
	result := Sum(2,2,3)

	if result != 6 {
		t.Fail()
	}

	// statement akan dieksekusi
	fmt.Println("TestFailSum dieksekusi")
}

// t.FailNow()
func TestFailNow(t *testing.T) {
	result := Sum(2,2,3)

	if result != 6 {
		t.FailNow()
	}

	// statement tidak di eksekusi
	fmt.Println("TestFailNow dieksekusi")
}

// t.Error()
func TestError(t *testing.T) {
	result := Sum(2,2,3)

	if result != 6 {
		t.Error("Expect = 6")
	}

	// statement di eksekusi
	fmt.Println("TestError dieksekusi")
}

// t.Fatal()
func TestFatal(t *testing.T) {
	result := Sum(2,2,3)

	if result != 6 {
		t.Fatal("Expect = 6")
	}

	// statement tidak di eksekusi
	fmt.Println("TestFatal dieksekusi")
}