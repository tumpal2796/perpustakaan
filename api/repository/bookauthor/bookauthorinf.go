package bookauthor

import (
	"context"

	"github.com/tokopedia/perpustakaan/api/model"
)

type BookAuthorInf interface {
	GetBookAuthor(ctx context.Context, bookIds []int64) (map[int64][]model.BooksAuthor, error)
}
