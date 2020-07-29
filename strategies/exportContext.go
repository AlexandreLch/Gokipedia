package strategies

import (
	"fmt"
	"gokipedia/models"
)

type Context struct {
	strategy ExportStrategy
}

func (c *Context) SetExportStrategy(e ExportStrategy) {
	c.strategy = e
}

func (c *Context) Export(articles []*models.Article) (*ArticleExport, error) {
	file, err := c.strategy.export(c, articles)
	if err != nil {
		return nil, fmt.Errorf("couldn't export: %v", err)
	}

	return file, nil
}
