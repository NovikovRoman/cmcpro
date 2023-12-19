package types

import "time"

type GlobalMetricsQuotesHistorical struct {
	Quotes []*GlobalMetricsQuote `json:"quotes,omitempty"`
}
type GlobalMetricsQuote struct {
	Timestamp    time.Time                          `json:"timestamp,omitempty"`
	BtcDominance float32                            `json:"btc_dominance,omitempty"`
	Quote        map[string]GlobalMetricsQuoteTotal `json:"quote,omitempty"`
}

type GlobalMetricsQuoteTotal struct {
	TotalMarketCap float64   `json:"total_market_cap,omitempty"`
	TotalVolume24h float64   `json:"total_volume_24h,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
}

type GlobalMetricsQuotesLatest struct {
	BtcDominance           float32                            `json:"btc_dominance,omitempty"`
	EthDominance           float32                            `json:"eth_dominance,omitempty"`
	ActiveCryptocurrencies uint                               `json:"active_cryptocurrencies,omitempty"`
	ActiveMarketPairs      uint                               `json:"active_market_pairs,omitempty"`
	ActiveExchanges        uint                               `json:"active_exchanges,omitempty"`
	LastUpdated            time.Time                          `json:"last_updated,omitempty"`
	Quote                  map[string]GlobalMetricsQuoteTotal `json:"quote,omitempty"`
}
