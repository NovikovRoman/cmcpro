package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_ExchangeQuotesLatestByID(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ExchangeQuotesLatestByID(
		[]uint{270}, NewConvertByCodes("RUB", "EUR"), // RUB, EUR
	)

	require.Equal(t, res["270"].Slug, "binance")
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.ExchangeQuotesLatestByID(
		[]uint{270}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["270"].Slug, "binance")
	require.NotNil(t, res["270"].Quote["RUB"])
	require.NotNil(t, res["270"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.ExchangeQuotesLatestByID(
		[]uint{270}, nil,
	)

	require.Equal(t, res["270"].Slug, "binance")
	require.NotNil(t, res["270"].Quote["USD"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.ExchangeQuotesLatestByID(
		[]uint{1000000000}, nil,
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_ExchangeQuotesLatestBySlug(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.ExchangeQuotesLatestBySlug(
		[]string{"bittrex"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["bittrex"].Slug, "bittrex")
	require.NotNil(t, res["bittrex"].Quote["RUB"])
	require.NotNil(t, res["bittrex"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.ExchangeQuotesLatestBySlug(
		[]string{"---"}, nil,
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_ExchangeQuotesHistoricalByID(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	timeStart := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 6))
	timeEnd := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 4))
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := c.ExchangeQuotesHistoricalByID(
		22, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 22)
	require.Equal(t, res.Name, "Bittrex")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = c.ExchangeQuotesHistoricalByID(
		22, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 22)
	require.Equal(t, res.Name, "Bittrex")
	require.True(t, len(res.Quotes) <= 5)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)
}

func TestClient_ExchangeQuotesHistoricalBySymbol(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	timeStart := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 6))
	timeEnd := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 4))
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := c.ExchangeQuotesHistoricalBySlug(
		"bittrex", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 22)
	require.Equal(t, res.Name, "Bittrex")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = c.ExchangeQuotesHistoricalBySlug(
		"bittrex", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 22)
	require.Equal(t, res.Name, "Bittrex")
	require.True(t, len(res.Quotes) <= 5)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)
}
