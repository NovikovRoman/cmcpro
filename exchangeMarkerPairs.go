package cmcpro

import (
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
)

func (c *Client) ExchangeMarketPairByID(
	id uint,
	start uint,
	limit uint,
	converter Converter,
) (*types.ExchangeMarketPairsLatest, *types.Status, error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}

	return c.exchangeMarketPair(params, start, limit, converter)
}

func (c *Client) ExchangeMarketPairBySlug(
	slug string,
	start uint,
	limit uint,
	converter Converter,
) (*types.ExchangeMarketPairsLatest, *types.Status, error) {

	params := map[string]string{
		"slug": slug,
	}

	return c.exchangeMarketPair(params, start, limit, converter)
}

func (c *Client) exchangeMarketPair(
	params map[string]string,
	start uint,
	limit uint,
	converter Converter,
) (*types.ExchangeMarketPairsLatest, *types.Status, error) {

	req, err := c.createRequest("/exchange/market-pairs/latest")
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

	if limit > ExchangeMaxLimit {
		limit = ExchangeMaxLimit
	} else if limit == 0 {
		limit = 1
	}
	query.Add("limit", strconv.FormatUint(uint64(limit), 10))

	if converter != nil {
		converter.AddQuery(&query)
	}

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   types.ExchangeMarketPairsLatest `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return &respInfo.Data, &respInfo.Status, nil
}
