package cmcpro

import (
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"strings"
)

func (c *Client) ExchangeMap(active bool, start uint, limit uint) ([]*types.ExchangeMap, *types.Status, error) {
	return c.exchangeMap(map[string]string{}, active, start, limit)
}

func (c *Client) ExchangeMapBySlug(slug []string) ([]*types.ExchangeMap, *types.Status, error) {
	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}
	return c.exchangeMap(params, false, 0, 0)
}

func (c *Client) exchangeMap(params map[string]string, active bool, start uint, limit uint) ([]*types.ExchangeMap, *types.Status, error) {
	req, err := c.createRequest("/exchange/map")
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

	if limit > ExchangeMaxLimit {
		limit = ExchangeMaxLimit
	} else if limit == 0 {
		limit = 1
	}
	query.Add("limit", strconv.FormatUint(uint64(limit), 10))

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   []*types.ExchangeMap `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
