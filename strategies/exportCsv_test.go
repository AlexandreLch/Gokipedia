package strategies

import (
	"bytes"
	"encoding/csv"
	. "github.com/onsi/gomega"
	"gokipedia/models"
	"log"
	"testing"
	"time"
)

func TestExportCsv(t *testing.T) {
	g := NewGomegaWithT(t)

	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	defer writer.Flush()

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

	var data [][]string
	for _, article := range articles {
		var articleArr []string
		articleArr = append(articleArr, string(article.ID), article.Title, article.Authors,
			article.Header, article.CreatedOn.Format(time.RFC3339), article.UpdatedOn.Format(time.RFC3339))
		data = append(data, articleArr)
	}

	err := writer.WriteAll(data)
	if err != nil {
		log.Printf("cannot write to file: %s", err)
	}

	g.Expect(b.Bytes()).ShouldNot(BeNil(), "Articles should have transformed to Bytes for CSV")
}
