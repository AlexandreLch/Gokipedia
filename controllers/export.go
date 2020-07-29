package controllers

import (
	"fmt"
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

	b, err := context.Export(articles)
	if err != nil {
		log.Printf("couldn't export: %v", err)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename%s=", b.FileName))
	w.Header().Set("Content-Type", r.Header.Get(b.MimeType))
	w.Write(b.FileByte)

	return
}
