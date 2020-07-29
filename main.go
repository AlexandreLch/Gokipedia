package main

import (
	"gokipedia/database"
	"gokipedia/router"
	"log"
	"net/http"
)

const dwldPath = "./tmp"

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	log.Print("\nServer started on port " + port)

	newRouter.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
		http.FileServer(http.Dir(dwldPath))))

	err = http.ListenAndServe(":"+port, newRouter)
	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
