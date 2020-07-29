package strategies

import (
	"log"
)

type ExportPdf struct{}

func (e *ExportPdf) export(c *Context) {
	log.Printf("Export as PDF")
}
