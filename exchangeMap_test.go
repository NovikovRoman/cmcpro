package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ExchangeMap(t *testing.T) {
	res, status, _ := cTest.ExchangeMap(true, 1, 10)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.ExchangeMapBySlug([]string{"binance", "poloniex"})

	require.Len(t, res, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}
