package domain

import (
	"time"
	"github.com/shopspring/decimal"
)

type Price struct {
	ID         UUID
	AssetUUID  UUID
	Price      decimal.Decimal
	Volume     decimal.Decimal
	RecordedAt time.Time
}