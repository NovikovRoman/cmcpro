package cmcpro

import (
	"testing"

	"github.com/NovikovRoman/cmcpro/types"
	"github.com/stretchr/testify/require"
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

	var res2 map[string][]types.CryptocurrencyInfo
	res2, status, _ = cTest.CryptocurrencyInfoBySymbol(contextTest, []string{"BTC", "LTC"})

	require.Len(t, res2, 2)
	require.Greater(t, len(res2["BTC"]), 0)
	require.NotNil(t, len(res2["LTC"]), 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}
