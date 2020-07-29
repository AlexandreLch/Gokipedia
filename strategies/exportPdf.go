package strategies

import "fmt"

type ExportPdf struct {}

func(e *ExportPdf) export(c *Context){
	fmt.Printf("Export as PDF")
}