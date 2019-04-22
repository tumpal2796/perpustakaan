package publisher

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tokopedia/orderapp/common/utils/strings"
	"github.com/tokopedia/perpustakaan/api/model"
)

const (
	queryPublisher = `select  * from publishers`
)

var (
	PublisherRes PublisherInf
)

type PublisherImpl struct {
	PsqlConn *sqlx.DB
}

func NewPublisherRes(dbConn *sqlx.DB) {
	PublisherRes = &PublisherImpl{
		PsqlConn: dbConn,
	}
}

func (pr *PublisherImpl) GetPublishers(ctx context.Context, publisherIds []int64) (map[int64]model.Publisher, error) {
	result := make(map[int64]model.Publisher)
	var queryResult []model.Publisher
	where := fmt.Sprintf(`where publisher_id in (%s)`, strings.JoinInt2String(publisherIds, ","))
	query := fmt.Sprintf("%s %s", queryPublisher, where)
	err := pr.PsqlConn.SelectContext(ctx, &queryResult, query)
	if err != nil {
		return result, err
	}

	for _, publisher := range queryResult {
		result[publisher.PublisherID] = publisher
	}

	return result, nil
}
