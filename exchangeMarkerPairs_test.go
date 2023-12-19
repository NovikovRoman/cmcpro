package cmcpro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ExchangeMarketPair(t *testing.T) {
	res, status, _ := cTest.ExchangeMarketPairByID(
		contextTest, 270, 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res.Name, "Binance")
	assert.Equal(t, res.Slug, "binance")
	assert.True(t, len(res.MarketPairs) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.ExchangeMarketPairBySlug(
		contextTest, "binance", 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res.Name, "Binance")
	assert.Equal(t, res.Slug, "binance")
	assert.True(t, len(res.MarketPairs) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
