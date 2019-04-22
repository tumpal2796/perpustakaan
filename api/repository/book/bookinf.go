package book

import (
	"context"

	"github.com/tokopedia/perpustakaan/api/model"
)

type BookInf interface {
	GetFileteredBooksIds(ctx context.Context, where string, args ...interface{}) ([]int64, error)
	GetBooks(ctx context.Context, bookIds []int64) (map[int64]model.Book, error)
}
