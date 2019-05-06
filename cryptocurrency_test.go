package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_CryptocurrencyListingsLatest(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.CryptocurrencyListingsLatest(
		1, 10, "symbol", "asc",
		NewConvertByCodes("USD"), "",
	)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}

func TestClient_CryptocurrencyListingsHistorical(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	date := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 4))

	res, status, _ := c.CryptocurrencyListingsHistorical(
		date, 1, 10, "symbol", "asc",
		NewConvertByCodes("USD"), "",
	)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
