package types

import "time"

type PriceConversion struct {
	ID          uint                                `json:"id"`
	Symbol      string                              `json:"symbol"`
	Name        string                              `json:"name"`
	Amount      float64                             `json:"amount,omitempty"`
	LastUpdated time.Time                           `json:"last_updated,omitempty"`
	Quote       map[string]*PriceConversionCurrency `json:"quote,omitempty"`
}

type PriceConversionCurrency struct {
	Price       float32   `json:"price,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
