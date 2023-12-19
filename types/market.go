package types

import "time"

type MarketPair struct {
	MarketID        uint                  `json:"market_id,omitempty"`
	MarketPair      string                `json:"market_pair,omitempty"`
	MarketPairBase  *PairInfo             `json:"market_pair_base,omitempty"`
	MarketPairQuote *PairInfo             `json:"market_pair_quote,omitempty"`
	Quote           map[string]*QuotePair `json:"quote,omitempty"`
}

type PairInfo struct {
	CurrencyID     uint   `json:"currency_id,omitempty"`
	CurrencySymbol string `json:"currency_symbol,omitempty"`
	CurrencyType   string `json:"currency_type,omitempty"`
}

type QuotePair struct {
	Price          float64   `json:"price,omitempty"`
	Volume24h      float64   `json:"volume_24h,omitempty"`
	Volume24hBase  float64   `json:"volume_24h_base,omitempty"`
	Volume24hQuote float64   `json:"volume_24h_quote,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
}
