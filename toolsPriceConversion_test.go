package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ToolsPriceConversion(t *testing.T) {
	res, status, _ := cTest.ToolsPriceConversionByID(
		10, 1, NewConvertByCodes("RUB"), nil,
	)

	require.Equal(t, res.Name, "Bitcoin")
	require.Equal(t, res.Amount, 10.0)
	require.True(t, res.Quote["RUB"].Price > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.ToolsPriceConversionBySymbol(
		10, "BTC", NewConvertByCodes("RUB"), nil,
	)

	require.Equal(t, res.Name, "Bitcoin")
	require.Equal(t, res.Amount, 10.0)
	require.True(t, res.Quote["RUB"].Price > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}
