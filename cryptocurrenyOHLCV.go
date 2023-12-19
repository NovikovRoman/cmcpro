package cmcpro

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyOHLCVLatestByID(ctx context.Context, id []uint, converter Converter) (data map[string]types.CryptocurrencyOHLCVLatest, status types.Status, err error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}
	res := struct {
		Data   map[string]types.CryptocurrencyOHLCVLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyOHLCVLatest(ctx, params, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyOHLCVLatestBySymbol(ctx context.Context, symbol []string, converter Converter) (data map[string][]types.CryptocurrencyOHLCVLatest, status types.Status, err error) {

	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}
	res := struct {
		Data   map[string][]types.CryptocurrencyOHLCVLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyOHLCVLatest(ctx, params, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) cryptocurrencyOHLCVLatest(ctx context.Context, params map[string]string, converter Converter, res interface{}) (err error) {

	req, err := c.createCryptocurrencyOHLCVRequest(ctx, "/v2/cryptocurrency/ohlcv/latest", params, nil, "", converter)
	if err != nil {
		return
	}
	return c.exec(req, &res)
}

func (c *Client) CryptocurrencyOHLCVHistoricalByID(ctx context.Context, id uint, perioder Perioder, interval string, converter Converter) (data types.CryptocurrencyOHLCVHistorical, status types.Status, err error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}
	res := struct {
		Data   types.CryptocurrencyOHLCVHistorical `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyOHLCVHistorical(ctx, params, perioder, interval, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyOHLCVHistoricalBySymbol(ctx context.Context, symbol string,
	perioder Perioder, interval string, converter Converter) (data map[string][]types.CryptocurrencyOHLCVHistorical, status types.Status, err error) {

	params := map[string]string{
		"symbol": symbol,
	}
	res := struct {
		Data   map[string][]types.CryptocurrencyOHLCVHistorical `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyOHLCVHistorical(ctx, params, perioder, interval, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) cryptocurrencyOHLCVHistorical(ctx context.Context, params map[string]string, perioder Perioder, interval string, converter Converter, res interface{}) (err error) {

	req, err := c.createCryptocurrencyOHLCVRequest(ctx,
		"/v2/cryptocurrency/ohlcv/historical", params, perioder, interval, converter)
	if err != nil {
		return
	}
	return c.exec(req, &res)
}

func (c *Client) createCryptocurrencyOHLCVRequest(ctx context.Context, link string, params map[string]string, perioder Perioder, interval string, converter Converter) (*http.Request, error) {

	req, err := c.createRequest(ctx, link)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	if perioder != nil {
		perioder.AddQuery(&query, true)
	}

	if interval != "" {
		query.Add("interval", interval)
	}

	if converter != nil {
		converter.AddQuery(&query)
	}

	req.URL.RawQuery = query.Encode()

	return req, nil
}
