package strategies

import (
	"log"
)

type ExportXls struct{}

func (e *ExportXls) export(c *Context, data []byte) error {
	log.Printf("Export as Xls")

	return nil
}
