package cmcpro

import (
	"testing"

	"github.com/NovikovRoman/cmcpro/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CryptocurrencyMarketPair(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyMarketPairByID(
		contextTest, 1, 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res.Name, "Bitcoin")
	assert.Len(t, res.MarketPairs, 5)
	assert.True(t, res.NumMarketPairs > 0)
	require.NotNil(t, res.MarketPairs[0].Quote["RUB"])
	assert.NotNil(t, res.MarketPairs[0].Quote["exchange_reported"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	var res2 []types.CryptocurrencyMarketPairsLatest
	res2, status, _ = cTest.CryptocurrencyMarketPairBySymbol(
		contextTest, "BTC", 1, 5, NewConvertByCodes("RUB"),
	)

	assert.Equal(t, res2[0].Name, "Bitcoin")
	assert.Len(t, res2[0].MarketPairs, 5)
	assert.True(t, res2[0].NumMarketPairs > 0)
	assert.NotNil(t, res2[0].MarketPairs[0].Quote["RUB"])
	assert.NotNil(t, res2[0].MarketPairs[0].Quote["exchange_reported"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}
}
