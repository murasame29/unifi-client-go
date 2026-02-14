package sitemanager

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

type GetISPMetricsParams struct {
	Type           string
	BeginTimestamp string
	EndTimestamp   string
	Duration       string
}

func (c *Client) GetISPMetrics(ctx context.Context, params GetISPMetricsParams) (*types.SingleResponse[[]types.ISPMetrics], error) {
	if params.Type == "" {
		return nil, errors.New("Type parameter is required")
	}

	query := url.Values{}
	if params.BeginTimestamp != "" {
		query.Set("beginTimestamp", params.BeginTimestamp)
	}
	if params.EndTimestamp != "" {
		query.Set("endTimestamp", params.EndTimestamp)
	}
	if params.Duration != "" {
		query.Set("duration", params.Duration)
	}

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/isp-metrics/%s", params.Type), query)
	if err != nil {
		return nil, err
	}

	var result types.SingleResponse[[]types.ISPMetrics]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
