package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_CryptocurrencyMap(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.CryptocurrencyMap(true, 1, 10)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.CryptocurrencyMapBySymbol([]string{"BTC", "LTC"})

	require.Len(t, res, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
