package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Author struct {
	AuthorID   int64         `db:"author_id"`
	Name       string        `db:"name"`
	CreateTime time.Time     `db:"create_time"`
	CreateBy   int64         `db:"create_by"`
	UpdateBy   pq.NullTime   `db:"update_time"`
	UpdateTime sql.NullInt64 `db:"update_by"`
}
