package cmcpro

import (
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"strings"
)

func (c *Client) CryptocurrencyInfoByID(id []uint) (map[string]*types.CryptocurrencyInfo, *types.Status, error) {
	ids := make([]string, len(id))

	for k, v := range id {
		ids[k] = strconv.FormatUint(uint64(v), 10)
	}

	params := map[string]string{
		"id": strings.Join(ids, ","),
	}

	return c.cryptocurrencyInfo(params)
}

func (c *Client) CryptocurrencyInfoBySlug(slug []string) (map[string]*types.CryptocurrencyInfo, *types.Status, error) {
	params := map[string]string{
		"slug": strings.Join(slug, ","),
	}

	return c.cryptocurrencyInfo(params)
}

func (c *Client) CryptocurrencyInfoBySymbol(symbol []string) (map[string]*types.CryptocurrencyInfo, *types.Status, error) {
	params := map[string]string{
		"symbol": strings.Join(symbol, ","),
	}

	return c.cryptocurrencyInfo(params)
}

func (c *Client) cryptocurrencyInfo(params map[string]string) (map[string]*types.CryptocurrencyInfo, *types.Status, error) {
	req, err := c.createRequest("/cryptocurrency/info")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   map[string]*types.CryptocurrencyInfo `json:"data"`
		Status types.Status
	}{}

	if err := c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
