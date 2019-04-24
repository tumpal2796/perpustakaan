package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Author struct {
	AuthorID   int64         `db:"author_id" json:"author_id"`
	Name       string        `db:"name" json:"name"`
	CreateTime time.Time     `db:"create_time" json:"create_time"`
	CreateBy   int64         `db:"create_by" json:"create_by"`
	UpdateBy   pq.NullTime   `db:"update_time" json:"update_time"`
	UpdateTime sql.NullInt64 `db:"update_by" json:"update_by"`
}
