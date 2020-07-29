package strategies

type ExportStrategy interface {
	export(*Context, [][]string) (*ArticleExport, error)
}

type ArticleExport struct {
	FileName string
	MimeType string
	FileByte []byte
}
