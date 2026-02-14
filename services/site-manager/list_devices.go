package sitemanager

import (
	"context"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

type ListDevicesParams struct {
	HostIDs   []string
	Time      string
	PageSize  string
	NextToken string
}

func (c *Client) ListDevices(ctx context.Context, params *ListDevicesParams) (*types.PaginatedResponse[types.DeviceHost], error) {
	query := url.Values{}
	if params != nil {
		for _, hostID := range params.HostIDs {
			query.Add("hostIds[]", hostID)
		}
		if params.Time != "" {
			query.Set("time", params.Time)
		}
		if params.PageSize != "" {
			query.Set("pageSize", params.PageSize)
		}
		if params.NextToken != "" {
			query.Set("nextToken", params.NextToken)
		}
	}

	resp, err := c.client.Get(ctx, "/v1/devices", query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.DeviceHost]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
