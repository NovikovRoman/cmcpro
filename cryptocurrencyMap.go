package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"strings"
)

func (c *Client) CryptocurrencyMap(ctx context.Context, active bool, start uint, limit uint) ([]*types.CryptocurrencyMap, *types.Status, error) {
	return c.cryptocurrencyMap(ctx, map[string]string{}, active, start, limit)
}

func (c *Client) CryptocurrencyMapBySymbol(ctx context.Context, symbol []string) ([]*types.CryptocurrencyMap, *types.Status, error) {
	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}
	return c.cryptocurrencyMap(ctx, params, false, 0, 0)
}

func (c *Client) cryptocurrencyMap(ctx context.Context, params map[string]string, active bool, start uint, limit uint) ([]*types.CryptocurrencyMap, *types.Status, error) {
	req, err := c.createRequest(ctx, "/cryptocurrency/map")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	if active {
		query.Add("listing_status", "active")
	} else {
		query.Add("listing_status", "inactive")
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
		Data   []*types.CryptocurrencyMap `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
