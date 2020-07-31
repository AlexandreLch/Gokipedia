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
	sheet, err := createSheet(file, articles)
	if err != nil {
		return nil, fmt.Errorf("could not create any sheet: %v", err)
	}

	for _, article := range articles {
		hydrateRows(article, sheet, articles)
	}

	file.Write(&b)

	return &ArticleExport{
		FileName: "articles.xlsx",
		MimeType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		FileByte: b.Bytes(),
	}, nil
}

func createSheet(file *xlsx.File, articles []*models.Article) (*xlsx.Sheet, error) {
	sheet, err := file.AddSheet("dinosaurs")
	if err != nil {
		return nil, fmt.Errorf("could not add sheet: %v", err)
	}

	row := sheet.AddRow()
	row.AddCell().Value = "Id"
	row.AddCell().Value = "Title"
	row.AddCell().Value = "Header"
	row.AddCell().Value = "Created on"
	row.AddCell().Value = "Updated on"

	for _, article := range articles {

		if len(article.Sections) > 0 {
			for range article.Sections {
				row.AddCell().Value = "Id"
				row.AddCell().Value = "Title"
				row.AddCell().Value = "Paragraph"
				row.AddCell().Value = "Position"
				row.AddCell().Value = "Media"
				row.AddCell().Value = "Created On"
				row.AddCell().Value = "Updated On"
				row.AddCell().Value = "Parent Id"
			}
		}
	}

	return sheet, nil
}

func hydrateRows(article *models.Article, sheet *xlsx.Sheet, articles []*models.Article) {
	row2 := sheet.AddRow()
	row2.AddCell().Value = strconv.FormatUint(article.ID, 10)
	row2.AddCell().Value = article.Title
	row2.AddCell().Value = article.Header
	row2.AddCell().Value = article.CreatedOn.Format(time.RFC3339)

	for _, article := range articles {
		if len(article.Sections) > 0 {
			for _, section := range article.Sections {
				row2.AddCell().Value = strconv.FormatUint(section.ID, 10)
				row2.AddCell().Value = section.Title
				row2.AddCell().Value = section.Paragraph
				row2.AddCell().Value = strconv.FormatUint(section.Position, 10)
				row2.AddCell().Value = section.Media
				row2.AddCell().Value = section.CreatedOn.Format(time.RFC3339)
				row2.AddCell().Value = section.UpdatedOn.Format(time.RFC3339)
				row2.AddCell().Value = strconv.FormatUint(section.ParentID, 10)
			}
		}
	}
}
