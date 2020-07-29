package strategies

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"gokipedia/models"
	"time"
)

type ExportCsv struct{}

func (e *ExportCsv) export(c *Context, articles []*models.Article) (*ArticleExport, error) {

	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	defer writer.Flush()

	var data [][]string
	for _, article := range articles {
		var articleArr []string
		articleArr = append(articleArr, string(article.ID), article.Title, article.Authors,
			article.Header, article.CreatedOn.Format(time.RFC3339), article.UpdatedOn.Format(time.RFC3339))
		data = append(data, articleArr)
	}

	err := writer.WriteAll(data)
	if err != nil {
		return nil, fmt.Errorf("cannot write to file: %s", err)
	}

	return &ArticleExport{
		FileName: "articles.csv",
		MimeType: "text/csv",
		FileByte: b.Bytes(),
	}, nil
}
