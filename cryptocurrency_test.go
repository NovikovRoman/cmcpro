package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_CryptocurrencyListingsLatest(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyListingsLatest(
		1, 10, "symbol", "asc",
		NewConvertByCodes("USD"), "",
	)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}

func TestClient_CryptocurrencyListingsHistorical(t *testing.T) {
	date := time.Now().Add(-time.Hour * 24 * 7 * 4)

	res, status, _ := cTest.CryptocurrencyListingsHistorical(
		date, 1, 10, "symbol", "asc",
		NewConvertByCodes("USD"), "",
	)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)
}
