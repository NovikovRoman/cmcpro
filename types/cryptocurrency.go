package types

import "time"

type CryptocurrencyHistorical struct {
	BasicCoin
	CmcRank           int                            `json:"cmc_rank"`
	NumMarketPairs    uint                           `json:"num_market_pairs"`
	CirculatingSupply float64                        `json:"circulating_supply,omitempty"`
	TotalSupply       float64                        `json:"total_supply,omitempty"`
	MaxSupply         float64                        `json:"max_supply,omitempty"`
	LastUpdated       time.Time                      `json:"last_updated,omitempty"`
	Quote             map[string]CryptocurrencyQuote `json:"quote,omitempty"`
}

type CryptocurrencyLatest struct {
	CryptocurrencyHistorical
	Tags      []string  `json:"tags,omitempty"`
	Platform  *Platform `json:"platform,omitempty"`
	DateAdded time.Time `json:"date_added,omitempty"`
}

type CryptocurrencyQuote struct {
	Price            float32   `json:"price,omitempty"`
	Volume24h        float64   `json:"volume_24h,omitempty"`
	PercentChange1h  float32   `json:"percent_change_1h,omitempty"`
	PercentChange24h float32   `json:"percent_change_24h,omitempty"`
	PercentChange7d  float32   `json:"percent_change_7d,omitempty"`
	MarketCap        float64   `json:"market_cap,omitempty"`
	LastUpdated      time.Time `json:"last_updated,omitempty"`
}
