package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"strings"
)

func (c *Client) ExchangeInfoByID(ctx context.Context, id []uint) (map[string]*types.ExchangeInfo, *types.Status, error) {
	ids := make([]string, len(id))
	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	return c.exchangeInfo(ctx, params)
}

func (c *Client) ExchangeInfoBySlug(ctx context.Context, slug []string) (map[string]*types.ExchangeInfo, *types.Status, error) {
	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}

	return c.exchangeInfo(ctx, params)
}

func (c *Client) exchangeInfo(ctx context.Context, params map[string]string) (map[string]*types.ExchangeInfo, *types.Status, error) {
	req, err := c.createRequest(ctx, "/exchange/info")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   map[string]*types.ExchangeInfo `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
