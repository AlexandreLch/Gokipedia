package strategies

import (
	"log"
)

type ExportCsv struct{}

func (e *ExportCsv) export(c *Context) {
	log.Printf("Export as CSV")
}
