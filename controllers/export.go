package controllers

import(
	"net/http"
	"log"
	"gokipedia/strategies"
)

func ExportArticles(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("exportType")
	log.Printf("%+v", name)

	var strategyType strategies.ExportStrategy

	switch(name){
	case "csv":
		strategyType = &strategies.ExportCsv{}
	case "pdf":
		strategyType = &strategies.ExportPdf{}
	}

	context := &strategies.Context{}
	context.SetExportStrategy(strategyType)
	context.Export()

	for key, value := range r.PostForm{
		log.Printf("\n%s = %s\n", key, value)
	}

	return
}
