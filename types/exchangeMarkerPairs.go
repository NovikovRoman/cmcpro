package types

type ExchangeMarketPairsLatest struct {
	BasicExchange
	NumMarketPairs uint          `json:"num_market_pairs,omitempty"`
	MarketPairs    []*MarketPair `json:"market_pairs,omitempty"`
}
