package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_ExchangeMarketPair(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ExchangeMarketPairByID(
		22, 1, 5, NewConvertByCodes("RUB"),
	)

	require.Equal(t, res.Name, "Bittrex")
	require.Len(t, res.MarketPairs, 5)
	require.True(t, res.NumMarketPairs > 5)
	require.NotNil(t, res.MarketPairs[0].Quote["RUB"])
	require.NotNil(t, res.MarketPairs[0].Quote["exchange_reported"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.ExchangeMarketPairBySlug(
		"bittrex", 1, 5, NewConvertByCodes("RUB"),
	)

	require.Equal(t, res.Name, "Bittrex")
	require.Len(t, res.MarketPairs, 5)
	require.True(t, res.NumMarketPairs > 5)
	require.NotNil(t, res.MarketPairs[0].Quote["RUB"])
	require.NotNil(t, res.MarketPairs[0].Quote["exchange_reported"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
