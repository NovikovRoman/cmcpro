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
