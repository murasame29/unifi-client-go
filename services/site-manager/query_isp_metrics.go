package sitemanager

import (
	"context"
	"fmt"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

func (c *Client) QueryISPMetrics(ctx context.Context, metricType string, req types.ISPMetricsQueryRequest) (*types.ISPMetricsQueryResponse, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/isp-metrics/%s/query", metricType), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ISPMetricsQueryResponse
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
