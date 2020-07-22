package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", HelloServer)
	port := "8080"
	log.Print("\nServer started on port "+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
