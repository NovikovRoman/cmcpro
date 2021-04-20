package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_CryptocurrencyInfo(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyInfoByID(contextTest, []uint{1, 1027})

	require.Len(t, res, 2)
	require.NotNil(t, res["1"])
	require.NotNil(t, res["1027"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.CryptocurrencyInfoBySlug(contextTest, []string{"bitcoin", "litecoin"})

	require.Len(t, res, 2)
	require.NotNil(t, res["1"])
	require.NotNil(t, res["2"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.CryptocurrencyInfoBySymbol(contextTest, []string{"BTC", "LTC"})

	require.Len(t, res, 2)
	require.NotNil(t, res["BTC"])
	require.NotNil(t, res["LTC"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}
