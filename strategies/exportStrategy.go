package strategies

type ExportStrategy interface {
	export(c *Context, data []byte) error
}
