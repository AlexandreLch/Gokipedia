package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func RenderArticles(w http.ResponseWriter, r *http.Request) {
	articlesData := []string{"Un article", "Un autre article"}
	tpl := template.New("articles.html")   
	path, _ := os.Getwd()
	log.Print(path)
	log.Printf("%+v", box)
	articles, err := box.FindString("articles.html")
	if err != nil {
		log.Print(err)
		return
	}
	t, err := tpl.Parse(articles) // Parse template file.
	if err != nil {
		log.Print(err)
		return
	}
	err = t.Execute(w, articlesData) // merge.
	if err != nil {
		log.Print(err)
		return
	}
}
