package types

import "time"

type CryptocurrencyMarketQuotesHistorical struct {
	ID     uint                          `json:"id"`
	Name   string                        `json:"name"`
	Symbol string                        `json:"symbol"`
	Quotes []*CryptocurrencyMarketQuotes `json:"quotes,omitempty"`
}

type CryptocurrencyMarketQuotes struct {
	Timestamp time.Time                             `json:"timestamp"`
	Quote     map[string]*CryptocurrencyMarketQuote `json:"quote,omitempty"`
}

type CryptocurrencyMarketQuote struct {
	Price       float32   `json:"price,omitempty"`
	Volume24h   float64   `json:"volume_24h,omitempty"`
	MarketCap   float64   `json:"market_cap,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}

type CryptocurrencyMarketQuotesLatest struct {
	ID                uint                                        `json:"id,omitempty"`
	Name              string                                      `json:"name,omitempty"`
	Symbol            string                                      `json:"symbol,omitempty"`
	Slug              string                                      `json:"slug,omitempty"`
	CirculatingSupply float64                                     `json:"circulating_supply,omitempty"`
	TotalSupply       float64                                     `json:"total_supply,omitempty"`
	MaxSupply         float64                                     `json:"max_supply,omitempty"`
	DateAdded         time.Time                                   `json:"date_added,omitempty"`
	NumMarketPairs    uint                                        `json:"num_market_pairs,omitempty"`
	CmcRank           int                                         `json:"cmc_rank,omitempty"`
	LastUpdated       time.Time                                   `json:"last_updated,omitempty"`
	Tags              []string                                    `json:"tags,omitempty"`
	Platform          *Platform                                   `json:"platform,omitempty"`
	Quote             map[string]*CryptocurrencyMarketQuoteLatest `json:"quote,omitempty"`
}

type CryptocurrencyMarketQuoteLatest struct {
	CryptocurrencyMarketQuote
	PercentChange1h  float32   `json:"percent_change_1h,omitempty"`
	PercentChange24h float32   `json:"percent_change_24h,omitempty"`
	PercentChange7d  float32   `json:"percent_change_7d,omitempty"`
	LastUpdated      time.Time `json:"last_updated,omitempty"`
}
