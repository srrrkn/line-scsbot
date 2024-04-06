package main
import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("src")))
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
