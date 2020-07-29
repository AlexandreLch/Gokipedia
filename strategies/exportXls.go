package strategies

import (
	"log"
	"os"
)

type ExportXls struct{}

func (e *ExportXls) export(c *Context, data [][]string) (*os.File, error) {
	log.Printf("Export as Xls")

	return nil, nil
}
