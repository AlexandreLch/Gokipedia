package controllers

import(
	"net/http"
	"log"
)

func ExportArticles(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("exportType")
	log.Printf("%+v", name)

	// new exportContext
	// setStrategy(csv)

	for key, value := range r.PostForm{
		log.Printf("%s = %s\n", key, value)
	}

	return
}
