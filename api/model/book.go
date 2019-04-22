package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Book struct {
	BookID      int64         `db:"book_id" json:"book_id"`
	Title       string        `db:"title" json:"title"`
	Cover       string        `db:"cover" json:"cover"`
	ISBN        string        `db:"isbn" json:"isbn"`
	Overview    string        `db:"overview" json:"overview"`
	PublisherID int64         `db:"publisher_id" json:"publisher_id"`
	CreateTime  time.Time     `db:"create_time" json:"create_time"`
	CreateBy    int64         `db:"create_by" json:"create_by"`
	UpdateBy    pq.NullTime   `db:"update_time" json:"update_time"`
	UpdateTime  sql.NullInt64 `db:"update_by" json:"update_by"`
}
