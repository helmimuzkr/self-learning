package golang_os

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

/* https://pkg.go.dev/os?GOOS=linux#File */

// Create File
func TestCreateFile(t *testing.T) {
	// text
	// text := "Hello Helmi!"

	// create file
	file, err := os.Create("coba.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Setiap selesai membuat file, file akan otomatis terbuka. Jadi harus ditutup setelah digunakan.
	defer file.Close()
	
	fName := file.Name()
	fmt.Println(fName)
}

// Create Temporary FIle
func TestCreateTempFile(t *testing.T) {
	// Path untuk lokasi file
	path := "C://Users//muzkr//Documents//Latihan//golang//golang-os//src//"

	// Membuat temporary file
	file, err := os.CreateTemp(path, "temporary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lihat nama file
	fmt.Println(file.Name())
}

// Open File
func TestOpenFile(t *testing.T) {
	// Kalau ga pakai src/ maka akan error. file tidak ditemukan
	file, err := os.Open("src/coba.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(file.Name())
}

// Write File
func TestWriteFile(t *testing.T) {
	// Open filenya dengan method OpenFile agar bisa menggunakan perm os.O_RDRW atau Read-Write
	file, err := os.OpenFile("src/coba.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("file berhasil dibuka")

	// Buat data string untuk dijadikan sebagai isi file.
	// Data string harus dikonversi menjadi slice of bytes
	data := []byte("hallo nama saya Helmi!")

	// Tulis data tadi ke dalam file dengan method Write()
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File berhasil diupdate!")
}


// Read File
func TestReadFile(t *testing.T) {
	file, err := os.Open("src/coba.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("file berhasil dibuka")

	// Buat variable slice bytes untuk dijadikan sebagai wadah penyimpanan data file, agar nantinya bisa di baca dengan println. sama seperti method Scan
	data := make([]byte, 255)

	for {
		n, err := file.Read(data)
		// Ketika errornya selain io.EOF maka proses baca file akan berlanjut.
		// Apa itu io.EOF? Error io.EOF sendiri menandakan bahwa file yang sedang dibaca adalah baris terakhir isi atau end of file.
		if err != io.EOF {
			break
		}
		// Kalau jumlah data yang ada pada file ada 0, maka perulangan dihentikan.
		if n == 0 {
			break
		}
	}

	fmt.Println(string(data))
}

