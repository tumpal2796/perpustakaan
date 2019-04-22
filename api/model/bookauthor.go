package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type BooksAuthor struct {
	BookAuthorID int64         `db:"book_author_id"`
	BookID       int64         `db:"book_id"`
	AuthorID     int64         `db:"author_id"`
	CreateTime   time.Time     `db:"create_time"`
	CreateBy     int64         `db:"create_by"`
	UpdateBy     pq.NullTime   `db:"update_time"`
	UpdateTime   sql.NullInt64 `db:"update_by"`
}
