package cmcpro

import (
	"context"
	"fmt"
	"github.com/NovikovRoman/cmcpro/types"
	"strconv"
	"time"
)

func (c *Client) ToolsPriceConversionByID(ctx context.Context, amount float64, id uint, converter Converter, t *time.Time) (*types.PriceConversion, *types.Status, error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}

	return c.toolsPriceConversion(ctx, params, amount, converter, t)
}

func (c *Client) ToolsPriceConversionBySymbol(ctx context.Context, amount float64, symbol string, converter Converter, t *time.Time) (*types.PriceConversion, *types.Status, error) {

	params := map[string]string{
		"symbol": symbol,
	}

	return c.toolsPriceConversion(ctx, params, amount, converter, t)
}

func (c *Client) toolsPriceConversion(ctx context.Context, params map[string]string, amount float64, converter Converter, t *time.Time) (*types.PriceConversion, *types.Status, error) {

	req, err := c.createRequest(ctx, "/tools/price-conversion")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()

	for n, v := range params {
		query.Add(n, v)
	}

	query.Add("amount", fmt.Sprint(amount))

	if t != nil {
		query.Add("time", t.Format(time.RFC3339))
	}

	if converter != nil {
		converter.AddQuery(&query)
	}

	req.URL.RawQuery = query.Encode()

	respInfo := struct {
		Data   *types.PriceConversion `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return respInfo.Data, &respInfo.Status, nil
}
