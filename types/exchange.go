package types

import "time"

type ExchangeHistorical struct {
	BasicExchange
	CmcRank        int                                `json:"cmc_rank"`
	NumMarketPairs uint                               `json:"num_market_pairs"`
	Timestamp      time.Time                          `json:"timestamp,omitempty"`
	Quote          map[string]ExchangeQuoteHistorical `json:"quote,omitempty"`
}

type ExchangeQuoteHistorical struct {
	ExchangeQuote
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type ExchangeLatest struct {
	BasicExchange
	// -1 unknown
	NumMarketPairs int                            `json:"num_market_pairs"`
	LastUpdated    time.Time                      `json:"last_updated,omitempty"`
	Quote          map[string]ExchangeQuoteLatest `json:"quote,omitempty"`
}

type ExchangeQuoteLatest struct {
	ExchangeQuote
	Volume24hAdjusted float64 `json:"volume_24h_adjusted,omitempty"`
}

type ExchangeQuote struct {
	Volume24h        float64 `json:"volume_24h,omitempty"`
	Volume7d         float64 `json:"volume_7d,omitempty"`
	Volume30d        float64 `json:"volume_30d,omitempty"`
	PercentChange24h float32 `json:"percent_change_volume_24h,omitempty"`
	PercentChange7d  float32 `json:"percent_change_volume_7d,omitempty"`
	PercentChange30d float32 `json:"percent_change_volume_30d,omitempty"`
}

type BasicExchange struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
