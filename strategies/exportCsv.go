package strategies

import (
	"encoding/csv"
	"fmt"
	"os"
)

type ExportCsv struct{}

func (e *ExportCsv) export(c *Context, data []byte) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return fmt.Errorf("couldn't create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("couldn't write to file: %v", err)
	}

	return nil
}
