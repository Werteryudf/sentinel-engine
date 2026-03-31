package domain

import "time"

type Asset struct {
	ID        UUID
	Ticker    string
	Name      string
	AssetType string
	CreatedAt time.Time
}