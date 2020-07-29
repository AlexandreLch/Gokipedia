package strategies

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

type ExportCsv struct{}

func (e *ExportCsv) export(c *Context, data [][]string) (*ArticleExport, error) {

	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	defer writer.Flush()

	err := writer.WriteAll(data)
	if err != nil {
		return nil, fmt.Errorf("cannot write to file: %s", err)
	}

	return &ArticleExport{
		FileName: "articles.csv",
		MimeType: "text/csv",
		FileByte: b.Bytes(),
	}, nil
}
