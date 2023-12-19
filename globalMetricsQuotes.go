package cmcpro

import (
	"context"
	"net/http"

	"github.com/NovikovRoman/cmcpro/types"
)

func (c *Client) GlobalMetricsQuotesLatest(ctx context.Context, converter Converter) (data types.GlobalMetricsQuotesLatest, status types.Status, err error) {

	req, err := c.createGlobalMetricsQuotesRequest(ctx, "/v1/global-metrics/quotes/latest", nil, converter)
	if err != nil {
		return
	}

	res := struct {
		Data   types.GlobalMetricsQuotesLatest `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) GlobalMetricsQuotesHistorical(ctx context.Context, perioder Perioder, converter Converter) (data types.GlobalMetricsQuotesHistorical, status types.Status, err error) {

	req, err := c.createGlobalMetricsQuotesRequest(ctx, "/v1/global-metrics/quotes/historical", perioder, converter)
	if err != nil {
		return
	}

	res := struct {
		Data   types.GlobalMetricsQuotesHistorical `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &res); err != nil {
		return
	}
	return res.Data, res.Status, nil
}

func (c *Client) createGlobalMetricsQuotesRequest(ctx context.Context, link string, perioder Perioder, converter Converter) (*http.Request, error) {
	req, err := c.createRequest(ctx, link)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()

	if perioder != nil {
		perioder.AddQuery(&query, false)
	}

	if converter != nil {
		converter.AddQuery(&query)
	}

	req.URL.RawQuery = query.Encode()
	return req, nil
}
