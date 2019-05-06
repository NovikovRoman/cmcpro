package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_ExchangeMap(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ExchangeMap(true, 1, 10)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.ExchangeMapBySlug([]string{"binance", "poloniex"})

	require.Len(t, res, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
