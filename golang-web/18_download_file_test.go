package web

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.RequestURI()
	log.Println("get url : ", URL)
	
	// get filename
	filename := r.URL.Query().Get("file")
	log.Println("get "+ filename)
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	// Add content disposition dengan value attachment, digunakan agar otomatis mendownload file tanpa render
	// w.Header().Add("Content-Disposition", "attachment; filename=" + filename)
	// w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	// log.Println("Content-Disposition", "attachment; filename=\""+filename+"\""+"added")

	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")

	// Tulis file yang mau dikirim ke response body
	http.ServeFile(w, r, "./resource/img/"+filename)
	log.Println("./resource/img/"+filename)
}

func TestDownloadFile(t *testing.T) {
	// route
	http.HandleFunc("/", downloadHandler)
	// http.HandlerFunc(downloadHandler)

	// listen and serve
	addr := "localhost:8080" 
	log.Println("listening on server : ", addr)
	err := http.ListenAndServe(addr, nil)
	// err := http.ListenAndServe(addr, http.HandlerFunc(downloadHandler))
	if err != nil {
		log.Fatal(err)
		return
	}
	
}

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")
	log.Println(file)

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestPZN(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}