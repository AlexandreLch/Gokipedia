package controllers

import (
	"gokipedia/strategies"
	"net/http"
)

func ExportArticles(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("exportType")

	var strategyType strategies.ExportStrategy

	switch name {
	case "csv":
		strategyType = &strategies.ExportCsv{}
	case "pdf":
		strategyType = &strategies.ExportPdf{}
	}

	context := &strategies.Context{}
	context.SetExportStrategy(strategyType)

	context.Export()

	return
}
