package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Println("Server started on :80")
	log.Fatal(http.ListenAndServe("", nil))
}
