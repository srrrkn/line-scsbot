package main

import (
	"fmt"
	"log"
	"net/http"
)

func reflectReply(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Test!")
}

func handleRequests() {
	http.HandleFunc("/webhook", reflectReply)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	handleRequests()
}