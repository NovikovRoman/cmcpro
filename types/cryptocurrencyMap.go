package types

import (
	"time"
)

type CryptocurrencyMap struct {
	BasicCoin
	IsActive            ConvertibleBoolean `json:"is_active"`
	FirstHistoricalData time.Time          `json:"first_historical_data,omitempty"`
	LastHistoricalData  time.Time          `json:"last_historical_data,omitempty"`
	Platform            *Platform          `json:"platform,omitempty"`
}
