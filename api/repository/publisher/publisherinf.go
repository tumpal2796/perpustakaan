package publisher

import (
	"context"

	"github.com/tokopedia/perpustakaan/api/model"
)

type PublisherInf interface {
	GetPublishers(ctx context.Context, publisherIds []int64) (map[int64]model.Publisher, error)
}
