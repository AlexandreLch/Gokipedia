package strategies

import (
	"log"
)

type ExportXls struct{}

func (e *ExportXls) export(c *Context) {
	log.Printf("Export as Xls")
}
