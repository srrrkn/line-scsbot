package main

import (
	"fmt"
	"log"
	"net/http"
)

func reflectReply(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Test!")
	fmt.Println("test")
	fmt.Println(r.Body)
}

func handleRequests() {
	http.HandleFunc("/webhook", reflectReply)
	log.Fatal(http.ListenAndServeTLS(":443", "/ssl/letsencrypt-all.crt", "/ssl/letsencrypt.key", nil))
}

func main() {
	handleRequests()
}