package strategies

import (
	"fmt"
)

type Context struct {
	strategy ExportStrategy
}

func (c *Context) SetExportStrategy(e ExportStrategy) {
	c.strategy = e
}

func (c *Context) Export(data [][]string) (*ArticleExport, error) {
	file, err := c.strategy.export(c, data)
	if err != nil {
		return nil, fmt.Errorf("couldn't export: %v", err)
	}

	return file, nil
}
