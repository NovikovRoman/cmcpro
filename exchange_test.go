package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ExchangeListingsLatest(t *testing.T) {
	res, status, _ := cTest.ExchangeListingsLatest(
		contextTest, 1, 2, "name", "asc",
		NewConvertByCodes("USD"), "",
	)

	require.Len(t, res, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}
}

// This endpoint is not yet available. It is slated for release by Q3 2019.
func TestClient_ExchangeListingsHistorical(t *testing.T) {
	/*c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	date, _ := time.Parse("02.01.2006", "01.01.2019")

	res, status, _ := c.ExchangeListingsHistorical(
		date, 1, 10,"name", "asc",
		NewConvertByCodes("USD"),"",
	)

	require.Len(t, res, 10)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)*/
}
