package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//go:embed resource
var resourceDir embed.FS

// Karena ketika directory diembed ke dalam variabel, maka direcotry patternnya "directory"
// jadi harus masuk ke folder "directory" supaya bisa mencapai file di dalamnya
// misalnya "directory/"
// Dengan menggunakan method fs.Sub, kita bisa masuk 1 step ke dalam folder.
// "directory" <== before
// "direcotry/" <== after
var directory, _ = fs.Sub(resourceDir, "resource")

// Create template
// Kenapa buatnya di luar handler? biar proses parsing hanya dilakukan sekali saja saat program dijalankan
// lalu bisa dipakai tanpa parsing lagi.
// Parsing template makan waktu.
// "*.*" artinnya string pattern tersebut memilih semua nama file dan semua extention
var templates = template.Must(template.ParseFS(directory, "*.*"))

// Index Handler
func formUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Data
	data := map[string]interface{}{
		"Title": "Tambah User",
	}
	
	templates.ExecuteTemplate(w, "index.html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.MultipartForm)
	name := r.PostFormValue("name")

	// Parsing form data file
	// @method r.ParseMultipartForm(), untuk memparsing form file yang ada data filenya
	// @arg maxMemory, Pemanggilan method tersebut membuat file yang terupload disimpan 
	// sementara pada memory dengan alokasi adalah sesuai dengan maxMemory.
	// @return error
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}

	// Ambil data file yang sudah diparsing
	// Sebenarnya, dengan method r.FormFile sudah ada proses parsing di dalamnya
	// @method r.FormFile, mengambil data form file yang diupload.
	// @arg key, untuk mengambil data dari data dari mutltipart.File dengan key ("picture").
	// @return multipart.File, sebagai representasi file yang diupload
	// @return *multipart.fileHeader, sebagai Informasi dari file, seperti nama file, ukuran, dll.
	// @return error
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// Membuat new file di local yang nantinya akan digunakan untuk menyimpan data dari file yang diupload
	fileDestination, err := os.Create("./resource/img/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	defer fileDestination.Close()

	// Menyalin data dari file yang diupload ke file baru yang ada di filesystem
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"Name": name,
		"Picture": "/picture/" + fileHeader.Filename,
	}

	templates.ExecuteTemplate(w, "profile.html", data)
}

func pictureHandler(w http.ResponseWriter, r *http.Request) {
	// get url
	URL := r.URL.RequestURI()
	log.Println("get url : ", URL)
	// Split URL dengan separator "/", return []string
	sliceURL := strings.Split(URL, "/")
	// Ambil element terakhir/filename
	filename := sliceURL[len(sliceURL)-1]
	
/* 	// Masuk ke directory local "img" supaya bisa mencapai file target
	directory, _ = fs.Sub(directory, "img")
	// Baca file dan convert menjadi []byte, supaya bisa dijadikan string untuk ditampilkan
	localFile, err := fs.ReadFile(directory, filename)
	if err != nil {
		panic(err)
	} */	

	if filename != "" {
		http.ServeFile(w, r, "./resource/img/"+filename)
		log.Println("./resource/img/"+filename)
	} else {
		fmt.Fprint(w, "error")
	}
}


func TestUploadFile(t *testing.T) {
	// Route
	mux := http.NewServeMux()
	mux.HandleFunc("/", formUploadHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/picture/", pictureHandler)
	
	// ini akan page not found karena di dalam folder resource(directory) tidak ada file yang dicari.
	// jadi harus masuk 1 step dulu menggunakan method fs.Sub. 8080:/picture/ktp.jpeg <== error, before
	// Agar tidak error, url yang yang ditulis harus seperti dibawah ini 
	// 8080:/picture/img/ktp.jpeg <== pass, before
	// dan, data yang dipassing ke response pada handler Upload harus diubah menjadi /picture/img/ + fileHeader.Filename
	// mux.Handle("/picture/", http.StripPrefix("/picture",http.FileServer(http.FS(directory))))

	// Jika tidak ingin menampilkan directory "img"-nya, maka kita harus masuk 1 folder lagi dengan cara menggunakan method fs.Sub.
	// Argumen yang dikirim adalah directory parrent dan patternstring directory img. Hasilnya nanti "/img/" 
	// directory, _ = fs.Sub(directory, "img") // 8080:/picture/ktp.jpeg
	// mux.Handle("/picture/", http.StripPrefix("/picture",http.FileServer(http.FS(directory)))) // 8080:/picture/ktp.jpeg

	// address
	addr := "localhost:8080"
	log.Println("running server on: ", addr)

	// runing server
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		panic(err)
	}
} 


// data dummy
//go:embed resource/img/download.png
var picture []byte

func TestUploadFileTesting(t *testing.T) {
	// Membuat body untuk menampung data yang akan dikirimkan sebagai request
	// Kenapa menggunakan bytes.Buffer? karena struct itu mengimplementasi io.Reader dan io.Writer interface
	bodyRequest := new(bytes.Buffer)

	// Membuat multipart writer dengan argumen bodyRequest
	// Writer ini nantinya digunakan untuk menulis data ke dalam bodyRequest
	multipartWriter := multipart.NewWriter(bodyRequest)
	// WriteField(name, value) digunakan untuk membuat request/menulis form text biasa
	err := multipartWriter.WriteField("name", "Helmi") 
	if err != nil {
		panic(err)
	}
	// CreateFormFile(name, filename) digunakan untuk membuat request/menulis form file
	writer, err := multipartWriter.CreateFormFile("file", "ContohTestUpload.png")
	if err != nil {
		panic(err)
	}
	// Menuliskan data dummy ke dalam bodyRequest
	writer.Write(picture)
	// Setelah selesai menggunakan multipartWriter harus langsung ditutup supaya bisa langsung disave
	// Kalau tidak disave, nanti saat get request form file di uploadHandler tidak akan ditemukan filenya
	// Jangan menggunakan defer
	multipartWriter.Close()
	

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", bodyRequest)
	// Set content type
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	recorder := httptest.NewRecorder()

	uploadHandler(recorder, request)

	bodyResponse, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(bodyResponse))
}

func TestFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Eko Kurniawan Khannedy")
	file, err := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	if err != nil {
		panic(err)
	}
	file.Write(picture)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	uploadHandler(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}