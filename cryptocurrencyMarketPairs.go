package cmcpro

import (
	"context"
	"strconv"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyMarketPairByID(ctx context.Context, id uint, start uint,
	limit uint, converter Converter) (data types.CryptocurrencyMarketPairsLatest, status types.Status, err error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}
	res := struct {
		Data   types.CryptocurrencyMarketPairsLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyMarketPair(ctx, params, start, limit, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) CryptocurrencyMarketPairBySymbol(ctx context.Context, symbol string, start uint, limit uint, converter Converter) (data []types.CryptocurrencyMarketPairsLatest, status types.Status, err error) {

	params := map[string]string{
		"symbol": symbol,
	}
	res := struct {
		Data   []types.CryptocurrencyMarketPairsLatest `json:"data"`
		Status types.Status
	}{}
	if err = c.cryptocurrencyMarketPair(ctx, params, start, limit, converter, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) cryptocurrencyMarketPair(ctx context.Context, params map[string]string,
	start uint, limit uint, converter Converter, res interface{}) (err error) {

	req, err := c.createRequest(ctx, "/v2/cryptocurrency/market-pairs/latest")
	if err != nil {
		return
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

	if converter != nil {
		converter.AddQuery(&query)
	}

	req.URL.RawQuery = query.Encode()
	return c.exec(req, &res)
}
