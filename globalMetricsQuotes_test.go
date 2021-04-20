package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_GlobalMetricsQuotesLatest(t *testing.T) {
	res, status, _ := cTest.GlobalMetricsQuotesLatest(NewConvertByCodes("RUB"))

	require.True(t, res.ActiveMarketPairs > 10)
	require.True(t, res.Quote["RUB"].TotalMarketCap > 1000.0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

}

func TestClient_GlobalMetricsQuotesHistorical(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.GlobalMetricsQuotesHistorical(perioder, NewConvertByCodes("RUB"))

	require.True(t, len(res.Quotes) <= 14)
	require.True(t, res.Quotes[0].Quote["RUB"].TotalMarketCap > 1000.0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
