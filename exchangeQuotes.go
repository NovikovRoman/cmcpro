package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"strings"
)

func (c *Client) ExchangeQuotesLatestByID(ctx context.Context, id []uint, converter Converter) (map[string]*types.ExchangeMarketQuotesLatest, *types.Status, error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	return c.exchangeQuotesLatest(ctx, params, converter)
}

func (c *Client) ExchangeQuotesLatestBySlug(ctx context.Context, slug []string, converter Converter) (map[string]*types.ExchangeMarketQuotesLatest, *types.Status, error) {

	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}

	return c.exchangeQuotesLatest(ctx, params, converter)
}

func (c *Client) exchangeQuotesLatest(ctx context.Context, params map[string]string, converter Converter) (map[string]*types.ExchangeMarketQuotesLatest, *types.Status, error) {

	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/v1/exchange/quotes/latest", params, nil, "", converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   map[string]*types.ExchangeMarketQuotesLatest `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

func (c *Client) ExchangeQuotesHistoricalByID(ctx context.Context, id uint, perioder Perioder, interval string, converter Converter) (*types.ExchangeMarketQuotesHistorical, *types.Status, error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}

	return c.exchangeQuotesHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) ExchangeQuotesHistoricalBySlug(ctx context.Context, slug string, perioder Perioder, interval string, converter Converter) (*types.ExchangeMarketQuotesHistorical, *types.Status, error) {

	params := map[string]string{
		"slug": slug,
	}

	return c.exchangeQuotesHistorical(ctx, params, perioder, interval, converter)
}

func (c *Client) exchangeQuotesHistorical(ctx context.Context, params map[string]string, perioder Perioder, interval string, converter Converter) (*types.ExchangeMarketQuotesHistorical, *types.Status, error) {

	req, err := c.createRequestCryptocurrencyQuotes(ctx, "/v1/exchange/quotes/historical", params, perioder, interval, converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   *types.ExchangeMarketQuotesHistorical `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

/* func (c *Client) createRequestExchangeQuotes(ctx context.Context, link string, params map[string]string, perioder Perioder, interval string, converter Converter) (*http.Request, error) {

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
} */
