package main
import (
	"log"
	"net/http"
)

func main() {
	port := "80"
	http.Handle("/", http.FileServer(http.Dir("/root/files")))
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
