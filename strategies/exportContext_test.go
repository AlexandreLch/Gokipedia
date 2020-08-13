package strategies

import (
	. "github.com/onsi/gomega"
	"gokipedia/models"
	"testing"
	"log"
	"time"
)


func TestExport(t *testing.T){
	g := NewGomegaWithT(t)

	context := Context{}
	context.SetExportStrategy(&ExportCsv{})

	articles := []*models.Article{
		&models.Article{
			ID: 1,
			Title: "Test",
			Header: "Test",
			Authors: "Nous",
			CreatedOn: time.Now(),
			UpdatedOn:  time.Now(),
			Sections: nil,
		},
	}

	b, err := context.Export(articles)
	if err != nil {
		log.Printf("couldn't export: %v", err)
		return
	}

	g.Expect(b).ShouldNot(BeNil(), "Articles should have transformed to Bytes")
}
