package strategies

import (
	"fmt"
	"gokipedia/models"
)

//Context is the context for the strategy pattern
type Context struct {
	strategy ExportStrategy
}

//SetExportStrategy set the strategy chosen for export (either csv or xls
func (c *Context) SetExportStrategy(e ExportStrategy) {
	c.strategy = e
}

//Export exports articles to the chosen export type
func (c *Context) Export(articles []*models.Article) (*ArticleExport, error) {
	file, err := c.strategy.export(c, articles)
	if err != nil {
		return nil, fmt.Errorf("couldn't export: %v", err)
	}

	return file, nil
}
