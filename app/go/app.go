package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5175"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "file-name"}),
	)

	http.ListenAndServe(":3001", corsHandler(mux))
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request")

	fileName := r.Header.Get("file-name")
	fmt.Println("File name:", fileName)
	uploadsDir := "./uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		os.Mkdir(uploadsDir, 0755)
	}
	filePath := filepath.Join(uploadsDir, fileName)

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(string(body))); err != nil {
		f.Close() 
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", fileName)
}
