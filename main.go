package main

import (
	"gokipedia/app/router"
	"log"
	"net/http"
)

func main() {
	newRouter := router.NewRouter()
	port := "8080"
	log.Print("\nServer started on port " + port)
	err := http.ListenAndServe(":"+port, newRouter)
	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
