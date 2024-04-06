package main

import (
	"fmt"
	"log"
	"net/http"
)

func reflectReply(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Test!")
	fmt.Println(w)
}

func handleRequests() {
	http.HandleFunc("/webhook", reflectReply)
	log.Fatal(http.ListenAndServe(":80", nil))
	log.Fatal(http.ListenAndServeTLS(":443", "/ssl/letsencrypt.crt", "/ssl/letsencrypt.key", nil))
}

func main() {
	handleRequests()
}