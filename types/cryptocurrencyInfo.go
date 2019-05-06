package types

type CryptocurrencyInfo struct {
	BasicCoin
	Category string           `json:"category,omitempty"`
	Logo     string           `json:"logo,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
	Platform *Platform        `json:"platform,omitempty"`
	Urls     *UrlTypeResponse `json:"urls,omitempty"`
}

type UrlTypeResponse struct {
	Website      []string `json:"website,omitempty"`
	Explorer     []string `json:"explorer,omitempty"`
	SourceCode   []string `json:"source_code,omitempty"`
	MessageBoard []string `json:"message_board,omitempty"`
	Chat         []string `json:"chat,omitempty"`
	Announcement []string `json:"announcement,omitempty"`
	Reddit       []string `json:"reddit,omitempty"`
	Twitter      []string `json:"twitter,omitempty"`
}
