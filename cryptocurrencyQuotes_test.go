package cmcpro

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CryptocurrencyQuotesLatestByID(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyQuotesLatestByID(
		contextTest, []uint{1}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["1"].Slug, "bitcoin")
	require.NotNil(t, res["1"].Quote["RUB"])
	require.NotNil(t, res["1"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestByID(
		contextTest, []uint{1}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["1"].Slug, "bitcoin")
	require.NotNil(t, res["1"].Quote["RUB"])
	require.NotNil(t, res["1"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestByID(
		contextTest, []uint{1}, nil,
	)

	require.Equal(t, res["1"].Slug, "bitcoin")
	require.NotNil(t, res["1"].Quote["USD"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestByID(
		contextTest, []uint{1000000000}, nil,
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_CryptocurrencyQuotesLatestBySlug(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyQuotesLatestBySlug(
		contextTest, []string{"ethereum"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["1027"].Slug, "ethereum")
	require.NotNil(t, res["1027"].Quote["RUB"])
	require.NotNil(t, res["1027"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestBySlug(
		contextTest, []string{"---"}, nil,
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_CryptocurrencyQuotesLatestBySymbol(t *testing.T) {
	res, status, _ := cTest.CryptocurrencyQuotesLatestBySymbol(
		contextTest, []string{"LTC"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["LTC"][0].Slug, "litecoin")
	require.NotNil(t, res["LTC"][0].Quote["RUB"])
	require.NotNil(t, res["LTC"][0].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestBySymbol(
		contextTest, []string{"1-2-3"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, status.ErrorCode, 0)
	require.Len(t, res["1-2-3"], 0)
}

func TestClient_CryptocurrencyQuotesHistoricalByID(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyQuotesHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	assert.NotNil(t, res)
	assert.EqualValues(t, res["1"].ID, 1)
	assert.Equal(t, res["1"].Name, "Bitcoin")
	assert.True(t, len(res["1"].Quotes) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyQuotesHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	assert.EqualValues(t, res["1"].ID, 1)
	assert.Equal(t, res["1"].Name, "Bitcoin")
	assert.True(t, len(res["1"].Quotes) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}
}

func TestClient_CryptocurrencyQuotesHistoricalBySymbol(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyQuotesHistoricalBySymbol(
		contextTest, "BTC", perioder, "1d", NewConvertByCodes("USD"),
	)

	assert.NotNil(t, res)
	assert.EqualValues(t, res["BTC"][0].ID, 1)
	assert.Equal(t, res["BTC"][0].Name, "Bitcoin")
	assert.True(t, len(res["BTC"][0].Quotes) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyQuotesHistoricalBySymbol(
		contextTest, "BTC", perioder, "7d", NewConvertByCodes("RUB", "EUR"),
	)

	assert.EqualValues(t, res["BTC"][0].ID, 1)
	assert.Equal(t, res["BTC"][0].Name, "Bitcoin")
	assert.True(t, len(res["BTC"][0].Quotes) > 0)
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}
