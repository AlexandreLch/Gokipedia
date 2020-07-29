package controllers

import (
	"gokipedia/database"
	"gokipedia/models"
	"gokipedia/strategies"
	"log"
	"net/http"
	"time"
)

func ExportArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	name := r.PostFormValue("exportType")

	var strategyType strategies.ExportStrategy

	switch name {
	case "csv":
		strategyType = &strategies.ExportCsv{}
	case "xls":
		strategyType = &strategies.ExportXls{}
	}

	context := &strategies.Context{}
	context.SetExportStrategy(strategyType)

	articles, err := repository.GetArticles()
	if err != nil {
		log.Printf("couldn't get articles: %v", err)
		return
	}

	var data [][]string
	for _, article := range articles {
		var articleArr []string
		articleArr = append(articleArr, string(article.ID), article.Title, article.Authors,
			article.Header, article.CreatedOn.Format(time.RFC3339), article.UpdatedOn.Format(time.RFC3339))
		data = append(data, articleArr)
	}

	b, err := context.Export(data)
	if err != nil {
		log.Printf("couldn't export: %v", err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=result.csv")
	w.Header().Set("Content-Type", r.Header.Get(b.MimeType))
	w.Write(b.FileByte)

	return
}
