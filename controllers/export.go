package controllers

import (
	"gokipedia/database"
	"gokipedia/models"
	"gokipedia/strategies"
	"log"
	"net/http"
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

	var data []string
	for _, article := range articles {
		data = append(data, string(article))
	}

	context.Export([]string{articles})

	return
}
