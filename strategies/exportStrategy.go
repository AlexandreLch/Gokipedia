package strategies

type ExportStrategy interface {
	export(*Context, [][]string) (*ArticleExport, error)
}

type ArticleExport struct {
	MimeType string
	FileByte []byte
}
