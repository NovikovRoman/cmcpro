package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) CryptocurrencyListingsLatest(ctx context.Context, start uint, limit uint, sort string, sortDir string,
	converter Converter, cryptocurrencyType string,
) ([]*types.CryptocurrencyLatest, *types.Status, error) {

	req, err := c.createCryptocurrencyListingsRequest(ctx, "/cryptocurrency/listings/latest",
		map[string]string{}, start, limit, sort, sortDir, converter, cryptocurrencyType,
	)

	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   []*types.CryptocurrencyLatest `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

func (c *Client) CryptocurrencyListingsHistorical(ctx context.Context, date time.Time,
	start uint, limit uint, sort string, sortDir string, converter Converter, cryptocurrencyType string,
) ([]*types.CryptocurrencyHistorical, *types.Status, error) {

	params := map[string]string{
		"date": date.Format(time.RFC3339),
	}

	req, err := c.createCryptocurrencyListingsRequest(ctx, "/cryptocurrency/listings/historical",
		params, start, limit, sort, sortDir, converter, cryptocurrencyType,
	)

	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   []*types.CryptocurrencyHistorical `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}

func (c *Client) createCryptocurrencyListingsRequest(ctx context.Context, link string,
	params map[string]string, start uint, limit uint, sort string, sortDir string,
	converter Converter, cryptocurrencyType string,
) (*http.Request, error) {

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

	if limit > CryptocurrencyMaxLimit {
		limit = CryptocurrencyMaxLimit
	} else if limit == 0 {
		limit = 1
	}
	query.Add("limit", strconv.FormatUint(uint64(limit), 10))

	if converter != nil {
		converter.AddQuery(&query)
	}

	if sort != "" {
		query.Add("sort", sort)
		sort := "asc"

		if sortDir != "" {
			sort = sortDir
		}

		query.Add("sort_dir", sort)
	}

	if cryptocurrencyType == All || cryptocurrencyType == Coins || cryptocurrencyType == Tokens {
		query.Add("cryptocurrency_type", cryptocurrencyType)
	}

	req.URL.RawQuery = query.Encode()

	return req, nil
}
