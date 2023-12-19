package cmcpro

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CryptocurrencyOHLCVLatestByID(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyOHLCVLatestByID(
		contextTest, []uint{1}, NewConvertByCodes("RUB", "EUR"),
	)

	assert.Equal(t, res["1"].Name, "Bitcoin")
	assert.NotNil(t, res["1"].Quote["RUB"])
	assert.NotNil(t, res["1"].Quote["EUR"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyOHLCVLatestByID(
		contextTest, []uint{1}, nil,
	)

	assert.Equal(t, res["1"].Name, "Bitcoin")
	assert.NotNil(t, res["1"].Quote["USD"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.CryptocurrencyOHLCVLatestByID(
		contextTest, []uint{10000000000}, nil,
	)

	assert.Equal(t, status.ErrorCode, 0)
	assert.Len(t, res, 0)
}

func TestClient_CryptocurrencyOHLCVLatestBySymbol(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyOHLCVLatestBySymbol(
		contextTest, []string{"LTC"}, NewConvertByCodes("RUB", "EUR"),
	)

	assert.Equal(t, res["LTC"][0].Name, "Litecoin")
	assert.NotNil(t, res["LTC"][0].Quote["RUB"])
	assert.NotNil(t, res["LTC"][0].Quote["EUR"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyOHLCVLatestBySymbol(
		contextTest, []string{"1-2-3"}, NewConvertByCodes("RUB", "EUR"),
	)

	assert.Equal(t, status.ErrorCode, 0)
	assert.Len(t, res, 0)
}

func TestClient_CryptocurrencyOHLCVHistoricalByID(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 4)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyOHLCVHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyOHLCVHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}

func TestClient_CryptocurrencyOHLCVHistoricalBySymbol(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyOHLCVHistoricalBySymbol(
		contextTest, "BTC", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res["BTC"][0].ID, 1)
	require.Equal(t, res["BTC"][0].Name, "Bitcoin")
	require.True(t, len(res["BTC"][0].Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyOHLCVHistoricalBySymbol(
		contextTest, "BTC", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res["BTC"][0].ID, 1)
	require.Equal(t, res["BTC"][0].Name, "Bitcoin")
	require.True(t, len(res["BTC"][0].Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}
