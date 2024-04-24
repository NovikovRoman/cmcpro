package cmcpro

import (
	"context"
	"strconv"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) CryptocurrencyCategory(ctx context.Context, id string, start uint, limit uint, converter Converter) (*types.CryptocurrencyCategoryData, *types.Status, error) {
	return c.cryptocurrencyCategory(ctx, map[string]string{"id": id}, start, limit, converter)
}

func (c *Client) cryptocurrencyCategory(ctx context.Context, params map[string]string, start uint, limit uint, converter Converter) (*types.CryptocurrencyCategoryData, *types.Status, error) {
	req, err := c.createRequest(ctx, "/v1/cryptocurrency/category")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()
	if converter != nil {
		converter.AddQuery(&query)
	}

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
		Data   *types.CryptocurrencyCategoryData `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}
	return respInfo.Data, &respInfo.Status, nil
}
