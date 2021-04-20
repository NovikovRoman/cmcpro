package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) CryptocurrencyOHLCVLatestByID(ctx context.Context, id []uint, converter Converter) (map[string]*types.CryptocurrencyOHLCVLatest, *types.Status, error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	return c.cryptocurrencyOHLCVLatest(ctx, params, converter)
}

func (c *Client) CryptocurrencyOHLCVLatestBySymbol(ctx context.Context, symbol []string, converter Converter) (map[string]*types.CryptocurrencyOHLCVLatest, *types.Status, error) {

	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}

	return c.cryptocurrencyOHLCVLatest(ctx, params, converter)
}

func (c *Client) cryptocurrencyOHLCVLatest(ctx context.Context, params map[string]string, converter Converter) (map[string]*types.CryptocurrencyOHLCVLatest, *types.Status, error) {

	req, err := c.createCryptocurrencyOHLCVRequest(ctx, "/cryptocurrency/ohlcv/latest", params, nil, "", converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   map[string]*types.CryptocurrencyOHLCVLatest `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

func (c *Client) CryptocurrencyOHLCVHistoricalByID(ctx context.Context, id uint, perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyOHLCVHistorical, *types.Status, error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}

	return c.cryptocurrencyOHLCVHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) CryptocurrencyOHLCVHistoricalBySymbol(ctx context.Context, symbol string,
	perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyOHLCVHistorical, *types.Status, error) {

	params := map[string]string{
		"symbol": symbol,
	}

	return c.cryptocurrencyOHLCVHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) cryptocurrencyOHLCVHistorical(ctx context.Context, params map[string]string, perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyOHLCVHistorical, *types.Status, error) {

	req, err := c.createCryptocurrencyOHLCVRequest(ctx,
		"/cryptocurrency/ohlcv/historical", params, perioder, interval, converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   *types.CryptocurrencyOHLCVHistorical `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
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
