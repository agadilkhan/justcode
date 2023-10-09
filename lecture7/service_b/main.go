package main

import (
	"io"
	"log"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/response"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	w.Write(body)
}

func response(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from service b"))
}

func main() {
	http.HandleFunc("/request", request)
	http.HandleFunc("/response", response)
	log.Println(http.ListenAndServe(":8081", nil))
}
