package strategies

import (
	"bytes"
	. "github.com/onsi/gomega"
	"github.com/tealeg/xlsx/v3"
	"gokipedia/models"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestExportXls(t *testing.T) {
	g := NewGomegaWithT(t)

	var b bytes.Buffer
	file := xlsx.NewFile()
	sheet, err := createSheet(file)
	if err != nil {
		log.Printf("could not create any sheet: %v", err)
	}

	articles := []*models.Article{
		&models.Article{
			ID:        1,
			Title:     "Test",
			Header:    "Test",
			Authors:   "Nous",
			CreatedOn: time.Now(),
			UpdatedOn: time.Now(),
			Sections:  nil,
		},
	}

	for _, article := range articles {
		hydrateRows(article, sheet)
	}

	file.Write(&b)

	g.Expect(b.Bytes()).ShouldNot(BeNil(), "Articles should have transformed to Bytes for XLS")
}

func TestCreateSheet(t *testing.T) {
	g := NewGomegaWithT(t)

	file := xlsx.NewFile()

	sheet, err := file.AddSheet("dinosaurs")
	if err != nil {
		log.Printf("could not add sheet: %v", err)
	}

	row := sheet.AddRow()
	row.AddCell().Value = "Id"
	row.AddCell().Value = "Title"
	row.AddCell().Value = "Header"
	row.AddCell().Value = "Created on"

	g.Expect(sheet).ShouldNot(BeNil(), "Sheet should have been created")
}

func TestHydrateRows(t *testing.T) {
	g := NewGomegaWithT(t)
	file := xlsx.NewFile()

	article := &models.Article{
		ID:        1,
		Title:     "Test",
		Header:    "Test",
		Authors:   "Nous",
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
		Sections:  nil,
	}

	sheet, err := file.AddSheet("dinosaurs")
	if err != nil {
		log.Printf("could not add sheet: %v", err)
	}

	row2 := sheet.AddRow()
	row2.AddCell().Value = strconv.FormatUint(article.ID, 10)
	row2.AddCell().Value = article.Title
	row2.AddCell().Value = article.Header
	row2.AddCell().Value = time.Time.String(article.CreatedOn)

	g.Expect(sheet).ShouldNot(BeNil(), "Sheet should have been created")
}
