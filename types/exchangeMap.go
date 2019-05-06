package types

import "time"

type ExchangeMap struct {
	ID                  uint               `json:"id"`
	Name                string             `json:"name"`
	Slug                string             `json:"slug"`
	IsActive            ConvertibleBoolean `json:"is_active"`
	FirstHistoricalData time.Time          `json:"first_historical_data,omitempty"`
	LastHistoricalData  time.Time          `json:"last_historical_data,omitempty"`
}
