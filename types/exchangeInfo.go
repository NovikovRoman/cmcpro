package types

type ExchangeInfo struct {
	ID   uint                     `json:"id"`
	Name string                   `json:"name"`
	Slug string                   `json:"slug"`
	Logo string                   `json:"logo,omitempty"`
	Urls *ExchangeUrlTypeResponse `json:"urls,omitempty"`
}

type ExchangeUrlTypeResponse struct {
	Website []string `json:"website,omitempty"`
	Twitter []string `json:"twitter,omitempty"`
	Blog    []string `json:"blog,omitempty"`
	Chat    []string `json:"chat,omitempty"`
	Fee     []string `json:"fee,omitempty"`
}
