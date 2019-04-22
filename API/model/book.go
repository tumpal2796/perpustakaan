package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Book struct {
	BookID      int64         `db:"book_id"`
	Title       string        `db:"title"`
	Cover       string        `db:"cover"`
	ISBN        string        `db:"isbn"`
	Overview    string        `db:"overview"`
	PublisherID int64         `db:"publisher_id"`
	CreateTime  time.Time     `db:"create_time"`
	CreateBy    int64         `db:"create_by"`
	UpdateBy    pq.NullTime   `db:"update_time"`
	UpdateTime  sql.NullInt64 `db:"update_by"`
}
