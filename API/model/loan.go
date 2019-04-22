package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Loan struct {
	LoanID     int64         `db:"loan_id"`
	BookID     int64         `db:"book_id"`
	UserID     int64         `db:"user_id"`
	DateOfLoan time.Time     `date_of_loan`
	Duration   int           `db:"duration"`
	CreateTime time.Time     `db:"create_time"`
	CreateBy   int64         `db:"create_by"`
	UpdateBy   pq.NullTime   `db:"update_time"`
	UpdateTime sql.NullInt64 `db:"update_by"`
}
