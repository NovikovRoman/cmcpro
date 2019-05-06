package types

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      uint   `json:"elapsed"`
	CreditCount  uint   `json:"credit_count"`
}
