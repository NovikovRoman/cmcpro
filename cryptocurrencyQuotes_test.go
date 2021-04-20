package cmcpro

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

	require.Equal(t, res["LTC"].Slug, "litecoin")
	require.NotNil(t, res["LTC"].Quote["RUB"])
	require.NotNil(t, res["LTC"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	res, status, _ = cTest.CryptocurrencyQuotesLatestBySymbol(
		contextTest, []string{"---"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_CryptocurrencyQuotesHistoricalByID(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyQuotesHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.NotNil(t, res)
	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyQuotesHistoricalByID(
		contextTest, 1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 4)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}

func TestClient_CryptocurrencyQuotesHistoricalBySymbol(t *testing.T) {
	timeStart := time.Now().Add(-time.Hour * 24 * 7 * 6)
	timeEnd := time.Now().Add(-time.Hour * 24 * 7 * 2)
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := cTest.CryptocurrencyQuotesHistoricalBySymbol(
		contextTest, "BTC", perioder, "1d", NewConvertByCodes("USD"),
	)

	require.NotNil(t, res)
	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 1)
	}

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = cTest.CryptocurrencyQuotesHistoricalBySymbol(
		contextTest, "BTC", perioder, "7d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.Len(t, res.Quotes, 2)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	if prodTest {
		require.EqualValues(t, status.CreditCount, 2)
	}
}
