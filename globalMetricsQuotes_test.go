package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_GlobalMetricsQuotesLatest(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.GlobalMetricsQuotesLatest(NewConvertByCodes("RUB"))

	require.True(t, res.ActiveMarketPairs > 10)
	require.True(t, res.Quote["RUB"].TotalMarketCap > float64(1000.0))
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

}

func TestClient_GlobalMetricsQuotesHistorical(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	timeStart := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 6))
	timeEnd := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 2))
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := c.GlobalMetricsQuotesHistorical(perioder, NewConvertByCodes("RUB"))

	require.True(t, len(res.Quotes) <= 14)
	require.True(t, res.Quotes[0].Quote["RUB"].TotalMarketCap > float64(1000.0))
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
