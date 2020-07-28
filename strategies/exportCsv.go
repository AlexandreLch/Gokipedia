package strategies

import "fmt"

type ExportCsv struct {
	Name string
}

func(e *ExportCsv) export(){
	fmt.Printf("Export as CSV")
}