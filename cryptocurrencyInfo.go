package cmcpro

import (
	"context"
	"strconv"
	"strings"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyInfoByID(ctx context.Context, id []uint) (info map[string]types.CryptocurrencyInfo, status types.Status, err error) {
	ids := make([]string, len(id))

	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	res := struct {
		Data   map[string]types.CryptocurrencyInfo `json:"data"`
		Status types.Status
	}{}

	if err = c.cryptocurrencyInfo(ctx, params, &res); err != nil {
		return
	}
	info = res.Data
	status = res.Status
	return
}

func (c *Client) CryptocurrencyInfoBySlug(ctx context.Context, slug []string) (info map[string]types.CryptocurrencyInfo, status types.Status, err error) {
	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}

	res := struct {
		Data   map[string]types.CryptocurrencyInfo `json:"data"`
		Status types.Status
	}{}

	if err = c.cryptocurrencyInfo(ctx, params, &res); err != nil {
		return
	}
	info = res.Data
	status = res.Status
	return
}

func (c *Client) CryptocurrencyInfoBySymbol(ctx context.Context, symbol []string) (info map[string][]types.CryptocurrencyInfo, status types.Status, err error) {
	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}

	res := struct {
		Data   map[string][]types.CryptocurrencyInfo `json:"data"`
		Status types.Status
	}{}

	if err = c.cryptocurrencyInfo(ctx, params, &res); err != nil {
		return
	}
	info = res.Data
	status = res.Status
	return
}

func (c *Client) cryptocurrencyInfo(ctx context.Context, params map[string]string, res interface{}) (err error) {
	req, err := c.createRequest(ctx, "/v2/cryptocurrency/info")
	if err != nil {
		return
	}

	query := req.URL.Query()
	for n, v := range params {
		query.Add(n, v)
	}
	req.URL.RawQuery = query.Encode()
	return c.exec(req, &res)
}
