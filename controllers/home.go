package controllers

import (
	"github.com/gobuffalo/packr/v2"
	"gokipedia/models"
	"html/template"
	"log"
	"net/http"
)

var box = packr.New("templateBox", "../templates")

func RenderHome(w http.ResponseWriter, r *http.Request) {
	message := models.Message{Greeting: "Yo frr"}
	tpl := template.New("home.html") // Create a template.
	home, err := box.FindString("home.html")
	if err != nil {
		log.Print(err)
		return
	}
	t, err := tpl.Parse(home) // Parse template file.
	if err != nil {
		log.Print(err)
		return
	}
	err = t.Execute(w, message) // merge.
	if err != nil {
		log.Print(err)
		return
	}
}
