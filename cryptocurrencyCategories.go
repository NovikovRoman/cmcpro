package cmcpro

import (
	"context"
	"strconv"
	"strings"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyCategories(ctx context.Context, start uint, limit uint) ([]*types.CryptocurrencyCategory, *types.Status, error) {
	return c.cryptocurrencyCategories(ctx, map[string]string{}, start, limit)
}

func (c *Client) CryptocurrencyCategoriesByCoinID(ctx context.Context, id []uint) ([]*types.CryptocurrencyCategory, *types.Status, error) {

	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}
	return c.cryptocurrencyCategories(ctx, params, 0, 0)
}

func (c *Client) CryptocurrencyCategoriesBySymbol(ctx context.Context, symbol []string) ([]*types.CryptocurrencyCategory, *types.Status, error) {
	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}
	return c.cryptocurrencyCategories(ctx, params, 0, 0)
}

func (c *Client) cryptocurrencyCategories(ctx context.Context, params map[string]string, start uint, limit uint) ([]*types.CryptocurrencyCategory, *types.Status, error) {
	req, err := c.createRequest(ctx, "/v1/cryptocurrency/categories")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	if start == 0 {
		start = 1
	}
	query.Add("start", strconv.FormatUint(uint64(start), 10))

	if limit > CryptocurrencyMaxLimit {
		limit = CryptocurrencyMaxLimit
	} else if limit == 0 {
		limit = 1
	}
	query.Add("limit", strconv.FormatUint(uint64(limit), 10))

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   []*types.CryptocurrencyCategory `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
