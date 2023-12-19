package cmcpro

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyQuotesLatestByID(ctx context.Context, id []uint, converter Converter) (data map[string]types.CryptocurrencyMarketQuotesLatest, status types.Status, err error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}
	res := struct {
		Data   map[string]types.CryptocurrencyMarketQuotesLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyQuotesLatest(ctx, params, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyQuotesLatestBySlug(ctx context.Context, slug []string, converter Converter) (data map[string]types.CryptocurrencyMarketQuotesLatest, status types.Status, err error) {

	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}
	res := struct {
		Data   map[string]types.CryptocurrencyMarketQuotesLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyQuotesLatest(ctx, params, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyQuotesLatestBySymbol(ctx context.Context, symbol []string, converter Converter) (data map[string][]types.CryptocurrencyMarketQuotesLatest, status types.Status, err error) {

	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}
	res := struct {
		Data   map[string][]types.CryptocurrencyMarketQuotesLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyQuotesLatest(ctx, params, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) cryptocurrencyQuotesLatest(ctx context.Context, params map[string]string, converter Converter, res interface{}) (err error) {
	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/v2/cryptocurrency/quotes/latest",
		params, nil, "", converter)
	if err != nil {
		return
	}
	return c.exec(req, &res)
}

func (c *Client) CryptocurrencyQuotesHistoricalByID(ctx context.Context, id uint, perioder Perioder, interval string, converter Converter) (data map[string]types.CryptocurrencyMarketQuotesHistorical, status types.Status, err error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}
	res := struct {
		Data   map[string]types.CryptocurrencyMarketQuotesHistorical `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyQuotesHistorical(ctx, params, perioder, interval, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyQuotesHistoricalBySymbol(ctx context.Context, symbol string, perioder Perioder, interval string, converter Converter) (data map[string][]types.CryptocurrencyMarketQuotesHistorical, status types.Status, err error) {

	params := map[string]string{
		"symbol": symbol,
	}
	res := struct {
		Data   map[string][]types.CryptocurrencyMarketQuotesHistorical `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyQuotesHistorical(ctx, params, perioder, interval, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) cryptocurrencyQuotesHistorical(ctx context.Context, params map[string]string, perioder Perioder, interval string, converter Converter, res interface{}) (err error) {

	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/v3/cryptocurrency/quotes/historical",
		params, perioder, interval, converter)
	if err != nil {
		return
	}
	return c.exec(req, &res)
}

func (c *Client) createRequestCryptocurrencyQuotes(ctx context.Context, link string, params map[string]string, perioder Perioder, interval string, converter Converter) (*http.Request, error) {

	req, err := c.createRequest(ctx, link)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	if perioder != nil {
		perioder.AddQuery(&query, false)
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
