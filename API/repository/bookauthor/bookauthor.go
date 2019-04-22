package bookauthor

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/orderapp/common/utils/strings"
	"github.com/tokopedia/perpustakaan/API/model"
)

const (
	queryBookAuthor = `select * from book_author`
)

var (
	BookAuthorRes BookAuthorInf
)

type BookAuthorImpl struct {
	PsqlConn *sqlx.DB
}

func New(dbConn *sqlx.DB) {
	BookAuthorRes = &BookAuthorImpl{
		PsqlConn: dbConn,
	}
}

func (bar *BookAuthorImpl) GetBookAuthor(ctx context.Context, bookIds []int64) (map[int64][]model.BooksAuthor, error) {
	result := make(map[int64][]model.BooksAuthor)
	var queryResult []model.BooksAuthor
	where := fmt.Sprintf(`where book_id in (%s)`, strings.JoinInt2String(bookIds, ","))
	query := fmt.Sprintf("%s %s", queryBookAuthor, where)
	err := bar.PsqlConn.SelectContext(ctx, &queryResult, query)
	if err != nil {
		return result, err
	}

	for _, bookauthor := range queryResult {
		result[bookauthor.BookID] = append(result[bookauthor.BookID], bookauthor)
	}

	return result, nil
}
