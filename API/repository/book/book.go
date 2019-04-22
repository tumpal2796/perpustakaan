package book

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/orderapp/common/utils/strings"
	"github.com/tokopedia/perpustakaan/API/model"
)

const (
	queryGetOrderIds = `select * from books`
)

var (
	BookRes BookInf
)

type BookImpl struct {
	PsqlConn *sqlx.DB
}

func NewBookRes(dbConn *sqlx.DB) {
	BookRes = &BookImpl{
		PsqlConn: dbConn,
	}
}

func (br *BookImpl) GetFileteredBooksIds(ctx context.Context, where string, args ...interface{}) ([]int64, error) {
	var result []int64
	var books []model.Book
	query := fmt.Sprintf("%s %s", queryGetOrderIds, where)
	err := br.PsqlConn.SelectContext(ctx, &books, query, args...)
	if err != nil {
		return result, err
	}

	for _, book := range books {
		result = append(result, book.BookID)
	}

	return result, nil
}

func (br *BookImpl) GetBooks(ctx context.Context, bookIds []int64) (map[int64]model.Book, error) {
	result := make(map[int64]model.Book)
	var queryResult []model.Book
	where := fmt.Sprintf(`where book_id in (%s)`, strings.JoinInt2String(bookIds, ","))
	query := fmt.Sprintf("%s %s", queryGetOrderIds, where)

	err := br.PsqlConn.SelectContext(ctx, &queryResult, query)
	if err != nil {
		return result, err
	}

	for _, book := range queryResult {
		result[book.BookID] = book
	}

	return result, nil
}
