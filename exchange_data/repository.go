package exchange_data

import (
	"forex/models"
	"time"
)

type ExchangeDataRepository interface {
	Create(*models.ExchangeData) (int64, error)
	GetAvgDay(start time.Time, end time.Time) ([]*models.ExchangeAvgAggregate, error)
}
