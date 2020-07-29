package strategies

import (
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"gokipedia/models"
	"strconv"
	"time"
)

type ExportXls struct{}

func (e *ExportXls) export(c *Context, articles []*models.Article) (*ArticleExport, error) {
	var b bytes.Buffer
	file := xlsx.NewFile()
	sheet, err := createSheet(file)
	if err != nil {
		return nil, fmt.Errorf("could not create any sheet: %v", err)
	}

	for _, article := range articles {
		hydrateRows(article, sheet)
	}

	file.Write(&b)

	return &ArticleExport{
		FileName: "articles.xlsx",
		MimeType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		FileByte: b.Bytes(),
	}, nil
}

func createSheet(file *xlsx.File) (*xlsx.Sheet, error) {
	sheet, err := file.AddSheet("dinosaurs")
	if err != nil {
		return nil, fmt.Errorf("could not add sheet: %v", err)
	}

	row := sheet.AddRow()
	row.AddCell().Value = "Id"
	row.AddCell().Value = "Title"
	row.AddCell().Value = "Header"
	row.AddCell().Value = "Created on"

	return sheet, nil
}

func hydrateRows(article *models.Article, sheet *xlsx.Sheet) {
	row2 := sheet.AddRow()
	row2.AddCell().Value = strconv.FormatUint(article.ID, 10)
	row2.AddCell().Value = article.Title
	row2.AddCell().Value = article.Header
	row2.AddCell().Value = time.Time.String(article.CreatedOn)
}
