package helper

/* Catatan

1. Untuk menjalankan unit test menggunakan command => "go test", jika di dalam folder test.
	ketika di root folder menggunakan command => "go test ./folder_test"
2. kalau ingin melihat apa saja yang ditest, tambahkan flag -v (verbose)
	contoh :
		a. "go test ./test -v"
		b. "go test -v -run=nama_func_test golang-testing/test"

*/

import (
	"testing"
)

// function yang akan di test
func Sum(input ...int) int {
	result := 0

	for _, v := range input {
		result += v
	}

	return result
}
func Mul(input ...int) int {
	var result int = 1

	for _, v := range input {
		result *= v
	}

	return result
}

// contoh berhasil
func TestSum(t *testing.T) {
	result := Sum(1,2,3)

	if result != 6 {
		t.Errorf("Sum(1,2,3) = %d; want 6", result)
	}
}

// contoh gagal
func TestSum2(t *testing.T) {
	result := Sum(2,2,3)

	if result != 6 {
		t.Errorf("Sum(1,2,3) = %d; want 6", result)
	}
}