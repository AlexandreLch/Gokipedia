package strategies

import (
	"log"
)

type ExportXls struct{}

func (e *ExportXls) export(c *Context, data [][]string) (*ArticleExport, error) {
	log.Printf("Export as Xls")

	return nil, nil
}
