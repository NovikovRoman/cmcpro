package types

import "time"

type ExchangeMarketQuotesHistorical struct {
	ID     uint              `json:"id"`
	Name   string            `json:"name"`
	Slug   string            `json:"slug"`
	Quotes []*ExchangeQuotes `json:"quotes,omitempty"`
}

type ExchangeQuotes struct {
	Timestamp      time.Time                      `json:"timestamp,omitempty"`
	Quote          map[string]*ExchangeQuotesItem `json:"quote,omitempty"`
	NumMarketPairs uint                           `json:"num_market_pairs,omitempty"`
}

type ExchangeQuotesItem struct {
	Volume24h float64   `json:"volume_24h,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type ExchangeMarketQuotesLatest struct {
	ID             uint                                `json:"id"`
	Name           string                              `json:"name"`
	Slug           string                              `json:"slug"`
	NumMarketPairs uint                                `json:"num_market_pairs,omitempty"`
	LastUpdated    time.Time                           `json:"last_updated,omitempty"`
	Quote          map[string]*ExchangeQuoteItemLatest `json:"quote,omitempty"`
}

type ExchangeQuoteItemLatest struct {
	Volume24h         float64 `json:"volume_24h,omitempty"`
	Volume24hAdjusted float64 `json:"volume_24h_adjusted,omitempty"`
	Volume7d          float64 `json:"volume_7d,omitempty"`
	Volume30d         float64 `json:"volume_30d,omitempty"`
	PercentChange24h  float32 `json:"percent_change_24h,omitempty"`
	PercentChange7d   float32 `json:"percent_change_7d,omitempty"`
	PercentChange30d  float32 `json:"percent_change_30d,omitempty"`
}
