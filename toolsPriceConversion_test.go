package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_ToolsPriceConversion(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ToolsPriceConversionByID(
		10, 1, NewConvertByCodes("RUB"), nil,
	)

	require.Equal(t, res.Name, "Bitcoin")
	require.Equal(t, res.Amount, float64(10.0))
	require.True(t, res.Quote["RUB"].Price > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.ToolsPriceConversionBySymbol(
		10, "BTC", NewConvertByCodes("RUB"), nil,
	)

	require.Equal(t, res.Name, "Bitcoin")
	require.Equal(t, res.Amount, float64(10.0))
	require.True(t, res.Quote["RUB"].Price > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
