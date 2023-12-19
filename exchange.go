package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
	"strconv"
)

func (c *Client) ExchangeListingsLatest(ctx context.Context, start uint,
	limit uint, sort string, sortDir string, converter Converter, marketType string) (data []types.ExchangeLatest, status types.Status, err error) {

	req, err := c.createExchangeListingsRequest(ctx, "/v1/exchange/listings/latest",
		map[string]string{}, start, limit, sort, sortDir, converter, marketType)

	if err != nil {
		return 
	}

	respInfo := struct {
		Data   []types.ExchangeLatest `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return
	}

	return respInfo.Data, respInfo.Status, nil
}

func (c *Client) createExchangeListingsRequest(ctx context.Context, link string, params map[string]string, start uint, limit uint, sort string, sortDir string, converter Converter, marketType string) (*http.Request, error) {

	req, err := c.createRequest(ctx, link)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
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

	if converter != nil {
		converter.AddQuery(&query)
	}

	if sort != "" {
		query.Add("sort", sort)
		sort = "asc"

		if sortDir != "" {
			sort = sortDir
		}

		query.Add("sort_dir", sort)
	}

	if marketType == All || marketType == NoFees || marketType == Fees {
		query.Add("market_type", marketType)
	}

	req.URL.RawQuery = query.Encode()

	return req, nil
}
