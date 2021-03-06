package author

import (
	"context"

	"github.com/tokopedia/perpustakaan/api/model"
)

type AuthorInf interface {
	GetAuthor(ctx context.Context, authorIds []int64) (map[int64]model.Author, error)
}
