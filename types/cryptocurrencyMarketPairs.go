package types

type CryptocurrencyMarketPairsLatest struct {
	ID             uint                        `json:"id"`
	Name           string                      `json:"name"`
	Symbol         string                      `json:"symbol"`
	NumMarketPairs uint                        `json:"num_market_pairs,omitempty"`
	MarketPairs    []*CryptocurrencyMarketPair `json:"market_pairs,omitempty"`
}

type CryptocurrencyMarketPair struct {
	MarketPair
	Exchange *BasicExchange
}
