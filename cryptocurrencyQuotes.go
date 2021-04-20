package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) CryptocurrencyQuotesLatestByID(ctx context.Context, id []uint, converter Converter) (map[string]*types.CryptocurrencyMarketQuotesLatest, *types.Status, error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	return c.cryptocurrencyQuotesLatest(ctx, params, converter)
}

func (c *Client) CryptocurrencyQuotesLatestBySlug(ctx context.Context, slug []string, converter Converter) (map[string]*types.CryptocurrencyMarketQuotesLatest, *types.Status, error) {

	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}

	return c.cryptocurrencyQuotesLatest(ctx, params, converter)
}

func (c *Client) CryptocurrencyQuotesLatestBySymbol(ctx context.Context, symbol []string, converter Converter) (map[string]*types.CryptocurrencyMarketQuotesLatest, *types.Status, error) {

	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}

	return c.cryptocurrencyQuotesLatest(ctx, params, converter)
}

func (c *Client) cryptocurrencyQuotesLatest(ctx context.Context, params map[string]string, converter Converter) (map[string]*types.CryptocurrencyMarketQuotesLatest, *types.Status, error) {

	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/cryptocurrency/quotes/latest",
		params, nil, "", converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   map[string]*types.CryptocurrencyMarketQuotesLatest `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

func (c *Client) CryptocurrencyQuotesHistoricalByID(ctx context.Context, id uint, perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyMarketQuotesHistorical, *types.Status, error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}

	return c.cryptocurrencyQuotesHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) CryptocurrencyQuotesHistoricalBySymbol(ctx context.Context, symbol string, perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyMarketQuotesHistorical, *types.Status, error) {

	params := map[string]string{
		"symbol": symbol,
	}

	return c.cryptocurrencyQuotesHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) cryptocurrencyQuotesHistorical(ctx context.Context, params map[string]string, perioder Perioder, interval string, converter Converter) (*types.CryptocurrencyMarketQuotesHistorical, *types.Status, error) {

	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/cryptocurrency/quotes/historical",
		params, perioder, interval, converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   *types.CryptocurrencyMarketQuotesHistorical `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
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
