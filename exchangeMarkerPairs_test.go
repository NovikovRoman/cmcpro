package cmcpro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ExchangeMarketPair(t *testing.T) {
	res, status, _ := cTest.ExchangeMarketPairByID(
		contextTest, 22, 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res.Name, "Bittrex Global")
	assert.Equal(t, res.Slug, "bittrex")
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.ExchangeMarketPairBySlug(
		contextTest, "bittrex", 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res.Name, "Bittrex Global")
	assert.Equal(t, res.Slug, "bittrex")
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
