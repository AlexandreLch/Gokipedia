package strategies

import "gokipedia/models"

//ExportStrategy interface for the strategy pattern
type ExportStrategy interface {
	export(*Context, []*models.Article) (*ArticleExport, error)
}

//ArticleExport DTO to export article
type ArticleExport struct {
	FileName string
	MimeType string
	FileByte []byte
}
