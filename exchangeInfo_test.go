package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_ExchangeInfo(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ExchangeInfoByID([]uint{270, 16})

	require.Len(t, res, 2)
	require.NotNil(t, res["270"])
	require.NotNil(t, res["16"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.ExchangeInfoBySlug([]string{"binance", "poloniex"})

	require.Len(t, res, 2)
	require.NotNil(t, res["binance"])
	require.NotNil(t, res["poloniex"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
