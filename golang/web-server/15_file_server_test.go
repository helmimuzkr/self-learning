package web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	// buat directory
	directory := http.Dir("./resource")

	// buat fileserver dan masukkan directorynya
	// Method http.FileServer akan mengembalikan http.Handler
	// yang nantinya bisa di assign ke handler server
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

//go:embed resource
var resource embed.FS

func TestFileServerEmbed(t *testing.T) {
	// Method fs.Sub berguna untuk masuk ke dalam folder resource. Seperti cd resource
	// Kenapa harus masuk ke dalam folder resource dulu? 
	// Ketika embed folder resource, pathnya = ./resource 
	// sehingga butuh masuk kedalam 1 step supaya bisa mencapai file yang ada di dalam folder resource
	// Contoh : 
	// http://localhost:8080/file/resource/index.html <== salah, tanpa melalui proses Sub
	// http://localhost:8080/file/index.html <== benar.
	// jadi, nama folder resource tidak ditampilkan.
	directory, err := fs.Sub(resource, "resource")
	if err != nil {
		panic(err)
	}

	// Buat file server handler
	// Karena hasil dari fs.Sub adalah fs.Fs, jadinya harus di parsing ke http.FileSystem
	// dengan method http.Fs
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}