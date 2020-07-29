package strategies

import "gokipedia/models"

type ExportStrategy interface {
	export(*Context, []*models.Article) (*ArticleExport, error)
}

type ArticleExport struct {
	FileName string
	MimeType string
	FileByte []byte
}
