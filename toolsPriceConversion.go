package cmcpro

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) ToolsPriceConversionByID(ctx context.Context, amount float64, id uint, converter Converter, t *time.Time) (data types.PriceConversion, status types.Status, err error) {

	params := map[string]string{
		"id": strconv.FormatUint(uint64(id), 10),
	}
	res := struct {
		Data   types.PriceConversion `json:"data"`
		Status types.Status
	}{}
	if err = c.toolsPriceConversion(ctx, params, amount, converter, t, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) ToolsPriceConversionBySymbol(ctx context.Context, amount float64, symbol string, converter Converter, t *time.Time) (data []types.PriceConversion, status types.Status, err error) {

	params := map[string]string{
		"symbol": symbol,
	}

	res := struct {
		Data   []types.PriceConversion `json:"data"`
		Status types.Status
	}{}
	if err = c.toolsPriceConversion(ctx, params, amount, converter, t, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) toolsPriceConversion(ctx context.Context, params map[string]string, amount float64, converter Converter, t *time.Time, res interface{}) (err error) {

	req, err := c.createRequest(ctx, "/v2/tools/price-conversion")
	if err != nil {
		return
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
	return c.exec(req, &res)
}
