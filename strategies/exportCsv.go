package strategies

import "fmt"

type ExportCsv struct {}

func(e *ExportCsv) export(c *Context){
	fmt.Printf("Export as CSV")
}