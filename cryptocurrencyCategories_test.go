package cmcpro

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CryptocurrencyCategories(t *testing.T) {
	res, status, err := cTest.CryptocurrencyCategories(contextTest, 1, 10)
	require.Nil(t, err)

	assert.Len(t, res, 10)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	res, status, err = cTest.CryptocurrencyCategoriesBySymbol(contextTest, []string{"BTC", "LTC"})
	require.Nil(t, err)

	assert.Len(t, res, 1)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	res, status, err = cTest.CryptocurrencyCategoriesByCoinID(contextTest, []uint{1, 1027})
	require.Nil(t, err)

	assert.Len(t, res, 1)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
