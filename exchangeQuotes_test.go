package cmcpro

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_ExchangeQuotesLatestByID(t *testing.T) {
	res, status, _ := cTest.ExchangeQuotesLatestByID(
		contextTest, []uint{270}, NewConvertByCodes("RUB", "EUR"), // RUB, EUR
	)

	assert.Equal(t, res["270"].Slug, "binance")
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.ExchangeQuotesLatestByID(
		contextTest, []uint{270}, NewConvertByCodes("RUB", "EUR"),
	)

	assert.Equal(t, res["270"].Slug, "binance")
	assert.NotNil(t, res["270"].Quote["RUB"])
	assert.NotNil(t, res["270"].Quote["EUR"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.ExchangeQuotesLatestByID(
		contextTest, []uint{270}, nil,
	)

	assert.Equal(t, res["270"].Slug, "binance")
	assert.NotNil(t, res["270"].Quote["USD"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	res, status, _ = cTest.ExchangeQuotesLatestByID(
		contextTest, []uint{1000000000}, nil,
	)

	assert.Equal(t, status.ErrorCode, 400)
	assert.Len(t, res, 0)
}

func TestClient_ExchangeQuotesLatestBySlug(t *testing.T) {
	res, status, _ := cTest.ExchangeQuotesLatestBySlug(
		contextTest, []string{"binance"}, NewConvertByCodes("RUB", "EUR"),
	)

	assert.Equal(t, res["binance"].Slug, "binance")
	assert.NotNil(t, res["binance"].Quote["RUB"])
	assert.NotNil(t, res["binance"].Quote["EUR"])
	assert.Equal(t, status.ErrorCode, 0)
	assert.Equal(t, status.ErrorMessage, "")
	if prodTest {
		assert.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.ExchangeQuotesLatestBySlug(
		contextTest, []string{"---"}, nil,
	)

	assert.Equal(t, status.ErrorCode, 400)
	assert.Len(t, res, 0)
}

func TestClient_ExchangeQuotesHistoricalByID(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 4)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.ExchangeQuotesHistoricalByID(
		contextTest, 270, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 270)
	require.Equal(t, res.Name, "Binance")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.ExchangeQuotesHistoricalByID(
		contextTest, 270, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 270)
	require.Equal(t, res.Name, "Binance")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)
}

func TestClient_ExchangeQuotesHistoricalBySymbol(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 4)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.ExchangeQuotesHistoricalBySlug(
		contextTest, "binance", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 270)
	require.Equal(t, res.Name, "Binance")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.ExchangeQuotesHistoricalBySlug(
		contextTest, "binance", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 270)
	require.Equal(t, res.Name, "Binance")
	require.True(t, len(res.Quotes) > 0)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}
