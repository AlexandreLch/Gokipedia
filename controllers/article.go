package controllers

import (
	"github.com/gorilla/mux"
	"gokipedia/database"
	"gokipedia/helpers"
	"gokipedia/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func RenderArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	articlesData, err := repository.GetArticles()
	if err != nil {
		log.Printf("could not get article list: %v", err)
		return
	}

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

func RenderArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	muxVars := mux.Vars(r)

	strID := muxVars["id"]

	intID, err := helpers.ParseUInt64(strID)

	articleData, err := repository.GetArticleByID(intID)
	if err != nil {
		log.Printf("could not get article data: %v", err)
		return
	}

	articleSections, err := repository.GetArticleContentByID(intID)
	if err != nil {
		log.Printf("could not get article content: %v", err)
		return
	}

	if len(articleSections) > 0 {
		articleData.Sections = articleSections
	}

	tpl := template.New("article.html")
	path, _ := os.Getwd()
	log.Print(path)
	log.Printf("%+v", box)
	article, err := box.FindString("article.html")
	if err != nil {
		log.Print(err)
		return
	}
	t, err := tpl.Parse(article) // Parse template file.
	if err != nil {
		log.Print(err)
		return
	}
	err = t.Execute(w, articleData) // merge.
	if err != nil {
		log.Print(err)
		return
	}
}

func RenderArticleForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("new-article.html")
	path, _ := os.Getwd()
	log.Print(path)
	log.Printf("%+v", box)
	articleForm, err := box.FindString("new-article.html")
	if err != nil {
		log.Print(err)
		return
	}
	t, err := tpl.Parse(articleForm) // Parse template file.
	if err != nil {
		log.Print(err)
		return
	}
	err = t.Execute(w, nil) // merge.
	if err != nil {
		log.Print(err)
		return
	}
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	r.ParseForm()
	article := models.Article{
		Title:     r.Form["title"],
		Header:    r.Form["header"],
		Authors:   "",
		CreatedOn: time.Time{},
		UpdatedOn: time.Time{},
		Sections:  nil,
	}
}
