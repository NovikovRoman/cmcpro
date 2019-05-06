package types

type BasicCoin struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}
