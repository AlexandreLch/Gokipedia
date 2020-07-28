package strategies

import "fmt"

type ExportPdf struct {
	Name string
}

func(e *ExportPdf) export(){
	fmt.Printf("Export as CSV")
}