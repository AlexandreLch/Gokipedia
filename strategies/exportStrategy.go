package strategies

import "os"

type ExportStrategy interface {
	export(*Context, [][]string) (*os.File, error)
}
