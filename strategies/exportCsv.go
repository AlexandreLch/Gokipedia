package strategies

import (
	"encoding/csv"
	"fmt"
	"os"
)

type ExportCsv struct{}

func (e *ExportCsv) export(c *Context, data [][]string) (*os.File, error) {
	file, err := os.Create("./tmp/result.csv")
	if err != nil {
		return nil, fmt.Errorf("couldn't create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return nil, fmt.Errorf("cannot write to file: %s", err)
		}
	}

	return file, nil
}
