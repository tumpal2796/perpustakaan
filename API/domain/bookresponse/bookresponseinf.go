package bookresponse

import (
	"context"

	"github.com/tokopedia/perpustakaan/API/model"
)

type BookResponseInf interface {
	GetBookResponse(ctx context.Context, filter model.Filter) ([]model.BookResponse, error)
}
