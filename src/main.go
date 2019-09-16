package main

import (
	"log"
	"net/http"

	"./hello"
)

func main() {
	http.HandleFunc("/hello", hello.NameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
