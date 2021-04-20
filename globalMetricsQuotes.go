package cmcpro

import (
	"context"
	"github.com/NovikovRoman/cmcpro/types"
	"net/http"
)

func (c *Client) GlobalMetricsQuotesLatest(ctx context.Context, converter Converter) (*types.GlobalMetricsQuotesLatest, *types.Status, error) {

	req, err := c.createGlobalMetricsQuotesRequest(ctx, "/global-metrics/quotes/latest", nil, converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   types.GlobalMetricsQuotesLatest `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return &respInfo.Data, &respInfo.Status, nil
}

func (c *Client) GlobalMetricsQuotesHistorical(ctx context.Context, perioder Perioder, converter Converter) (*types.GlobalMetricsQuotesHistorical, *types.Status, error) {

	req, err := c.createGlobalMetricsQuotesRequest(ctx, "/global-metrics/quotes/historical", perioder, converter)
	if err != nil {
		return nil, nil, err
	}

	respInfo := struct {
		Data   types.GlobalMetricsQuotesHistorical `json:"data"`
		Status types.Status
	}{}

	if err = c.exec(req, &respInfo); err != nil {
		return nil, nil, err
	}

	return &respInfo.Data, &respInfo.Status, nil
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
