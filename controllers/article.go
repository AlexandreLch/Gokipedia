package controllers

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gokipedia/database"
	"gokipedia/helpers"
	"gokipedia/models"
	"html/template"
	"log"
	"net/http"
)

var decoder = schema.NewDecoder()

func RenderArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	articlesData, err := repository.GetArticles()
	if err != nil {
		log.Printf("could not get article list: %v", err)
		return
	}

	tpl := template.New("articles.html")
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

	err := r.ParseForm()
	if err != nil {
		log.Printf("could not parse form: %v", err)
		return
	}

	var article models.Article

	err = decoder.Decode(&article, r.PostForm)
	if err != nil {
		log.Printf("could not decode form into article: %v", err)
		return
	}

	err = repository.SaveArticle(&article)
	if err != nil {
		log.Printf("could not save article: %v", err)
		return
	}

	log.Print("article saved")
}
