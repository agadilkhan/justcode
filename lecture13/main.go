package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fileName := path.Base(r.URL.Path)

	fileBytes, err := os.ReadFile(fmt.Sprintf("lecture13/internal/files/%s", fileName))
	if err != nil {
		log.Println(err)
		w.Write([]byte("not found this file"))

		return
	}
	w.Write(fileBytes)
}
