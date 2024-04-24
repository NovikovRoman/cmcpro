package types

type CryptocurrencyCategoryData struct {
	CryptocurrencyCategory
	Coins []CryptocurrencyLatest `json:"coins"`
}
