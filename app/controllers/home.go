package controllers

import (
	"fmt"
	"gokipedia/app/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	message := models.Message{Greeting: "Yo frr"}
	tpl := template.New("home.html")                   // Create a template.
	t, err := tpl.ParseFiles("go/templates/home.html") // Parse template file.
	if err != nil {
		log.Print(err)
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path) // for example /home/user
		return
	}
	err = t.Execute(w, message) // merge.
	if err != nil {
		log.Print(err)
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path) // for example /home/user
		return
	}
}
