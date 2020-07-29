package controllers

import (
	"fmt"
	"gokipedia/database"
	"gokipedia/models"
	"gokipedia/strategies"
	"io"
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

	var data [][]string
	for _, article := range articles {
		var articleArr []string
		articleArr = append(articleArr, article.Title, article.Authors, article.Header)
		data = append(data, articleArr)
	}

	//file, err := context.Export(data)
	//if err != nil {
	//	log.Printf("couldn't export: %v", err)
	//	return
	//}

	url := "http://localhost:8080/files/result.csv"
	client := http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=result.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("could not download file: %v", err)
		return
	}

	return
}
