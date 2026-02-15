package sitemanager

import (
	"context"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

func (c *Client) ListSDWANConfigs(ctx context.Context) (*types.SingleResponse[[]types.SDWANConfig], error) {
	resp, err := c.client.Get(ctx, "/v1/sd-wan-configs", nil)
	if err != nil {
		return nil, err
	}

	var result types.SingleResponse[[]types.SDWANConfig]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
