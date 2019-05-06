package types

import "time"

type CryptocurrencyOHLCVHistorical struct {
	ID     uint                         `json:"id"`
	Name   string                       `json:"name"`
	Symbol string                       `json:"symbol"`
	Quotes []*CryptocurrencyOHLCVQuotes `json:"quotes"`
}

type CryptocurrencyOHLCVQuotes struct {
	TimeOpen  time.Time                            `json:"time_open,omitempty"`
	TimeClose time.Time                            `json:"time_close,omitempty"`
	Quote     map[string]*CryptocurrencyOHLCVQuote `json:"quote,omitempty"`
}

type CryptocurrencyOHLCVQuote struct {
	OHLCV
	MarketCap float64   `json:"market_cap,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type OHLCV struct {
	Open   float32 `json:"open,omitempty"`
	High   float32 `json:"high,omitempty"`
	Low    float32 `json:"low,omitempty"`
	Close  float32 `json:"close,omitempty"`
	Volume float64 `json:"volume,omitempty"`
}

type CryptocurrencyOHLCVLatest struct {
	CryptocurrencyOHLCVHistorical
	LastUpdated time.Time                                  `json:"last_updated,omitempty"`
	TimeOpen    time.Time                                  `json:"time_open,omitempty"`
	TimeClose   time.Time                                  `json:"time_close,omitempty"`
	Quote       map[string]*CryptocurrencyOHLCVQuoteLatest `json:"quote,omitempty"`
}

type CryptocurrencyOHLCVQuoteLatest struct {
	OHLCV
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
