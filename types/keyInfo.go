package types

type KeyInfo struct {
	Plan struct {
		CreditLimitMonthly               int    `json:"credit_limit_monthly"`
		CreditLimitMonthlyReset          string `json:"credit_limit_monthly_reset"`
		CreditLimitMonthlyResetTimestamp string `json:"credit_limit_monthly_reset_timestamp"`
		RateLimitMinute                  int    `json:"rate_limit_minute"`
	} `json:"plan"`

	Usage struct {
		Minute struct {
			RequestsMade int `json:"requests_made"`
			RequestsLeft int `json:"requests_left"`
		} `json:"current_minute"`
		CurrentDay struct {
			CreditsUsed int `json:"credits_used"`
			CreditsLeft int `json:"credits_left"`
		} `json:"current_day"`
		CurrentMonth struct {
			CreditsUsed int `json:"credits_used"`
			CreditsLeft int `json:"credits_left"`
		} `json:"current_month"`
	} `json:"usage"`
}
