package cmcpro

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_CryptocurrencyOHLCVLatestByID(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	res, status, _ := c.CryptocurrencyOHLCVLatestByID(
		[]uint{1}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["1"].Name, "Bitcoin")
	require.NotNil(t, res["1"].Quote["RUB"])
	require.NotNil(t, res["1"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.CryptocurrencyOHLCVLatestByID(
		[]uint{1}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["1"].Name, "Bitcoin")
	require.NotNil(t, res["1"].Quote["RUB"])
	require.NotNil(t, res["1"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.CryptocurrencyOHLCVLatestByID(
		[]uint{1}, nil,
	)

	require.Equal(t, res["1"].Name, "Bitcoin")
	require.NotNil(t, res["1"].Quote["USD"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 1)

	res, status, _ = c.CryptocurrencyOHLCVLatestByID(
		[]uint{10000000000}, nil,
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_CryptocurrencyOHLCVLatestBySymbol(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)
	res, status, _ := c.CryptocurrencyOHLCVLatestBySymbol(
		[]string{"LTC"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, res["LTC"].Name, "Litecoin")
	require.NotNil(t, res["LTC"].Quote["RUB"])
	require.NotNil(t, res["LTC"].Quote["EUR"])
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	res, status, _ = c.CryptocurrencyOHLCVLatestBySymbol(
		[]string{"---"}, NewConvertByCodes("RUB", "EUR"),
	)

	require.Equal(t, status.ErrorCode, 400)
	require.Len(t, res, 0)
}

func TestClient_CryptocurrencyOHLCVHistoricalByID(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	timeStart := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 6))
	timeEnd := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 4))
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := c.CryptocurrencyOHLCVHistoricalByID(
		1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = c.CryptocurrencyOHLCVHistoricalByID(
		1, perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 5)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)
}

func TestClient_CryptocurrencyOHLCVHistoricalBySymbol(t *testing.T) {
	c, _ := New(os.Getenv("API-KEY"), false, "", Timeout)

	timeStart := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 6))
	timeEnd := time.Now().Add(- time.Duration(time.Hour * 24 * 7 * 2))
	perioder := NewPeriod(&timeStart, &timeEnd, 50)

	res, status, _ := c.CryptocurrencyOHLCVHistoricalBySymbol(
		"BTC", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 14)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)

	perioder = NewPeriod(&timeStart, nil, 5)
	res, status, _ = c.CryptocurrencyOHLCVHistoricalBySymbol(
		"BTC", perioder, "1d", NewConvertByCodes("RUB", "EUR"),
	)

	require.EqualValues(t, res.ID, 1)
	require.Equal(t, res.Name, "Bitcoin")
	require.True(t, len(res.Quotes) <= 5)
	require.Equal(t, status.ErrorCode, 0)
	require.Equal(t, status.ErrorMessage, "")
	require.EqualValues(t, status.CreditCount, 2)
}
