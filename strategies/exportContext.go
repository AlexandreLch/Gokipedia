package strategies

type Context struct {
	strategy ExportStrategy
}

func (c *Context) SetExportStrategy(e ExportStrategy) {
	c.strategy = e
}

func (c *Context) Export() {
	c.strategy.export(c)
}
