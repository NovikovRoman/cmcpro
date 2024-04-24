package types

import (
	"time"
)

type CryptocurrencyCategory struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	NumTokens       uint      `json:"num_tokens"`
	AvgPriceChange  float64   `json:"avg_price_change"`
	MarketCap       float64   `json:"market_cap"`
	MarketCapChange float64   `json:"market_cap_change"`
	Volume          float64   `json:"volume"`
	VolumeChange    float64   `json:"volume_change"`
	LastUpdated     time.Time `json:"last_updated"`
}
