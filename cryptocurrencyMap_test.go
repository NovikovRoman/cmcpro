package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_CryptocurrencyMap(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyMap(true, 1, 10)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.CryptocurrencyMapBySymbol([]string{"BTC", "LTC"})

	require.Len(t, res, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}
