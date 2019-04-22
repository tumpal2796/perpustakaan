package author

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/orderapp/common/utils/strings"
	"github.com/tokopedia/perpustakaan/api/model"
)

var (
	queryAuthor = `select * from authors`
)

var (
	AuthorRes AuthorInf
)

type AuthorImpl struct {
	PsqlConn *sqlx.DB
}

func NewAuthorRes(dbConn *sqlx.DB) {
	AuthorRes = &AuthorImpl{
		PsqlConn: dbConn,
	}
}

func (ar *AuthorImpl) GetAuthor(ctx context.Context, authorIds []int64) (map[int64]model.Author, error) {
	result := make(map[int64]model.Author)
	var queryResult []model.Author
	where := fmt.Sprintf(`where author_id in (%s)`, strings.JoinInt2String(authorIds, ","))
	query := fmt.Sprintf("%s %s", queryAuthor, where)

	err := ar.PsqlConn.SelectContext(ctx, &queryResult, query)
	if err != nil {
		return result, nil
	}

	for _, author := range queryResult {
		result[author.AuthorID] = author
	}

	return result, nil
}
