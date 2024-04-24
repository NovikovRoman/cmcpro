package cmcpro

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CryptocurrencyCategory(t *testing.T) {
	res, status, err := cTest.CryptocurrencyCategory(contextTest, "6051a82566fc1b42617d6dc6", 1, 10, nil)
	require.Nil(t, err)

	assert.Len(t, res.Coins, 10)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
