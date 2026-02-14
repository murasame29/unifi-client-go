package network

import (
	"context"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetApplicationInfo(ctx context.Context) (*types.ApplicationInfo, error) {
	resp, err := c.client.Get(ctx, "/v1/info", nil)
	if err != nil {
		return nil, err
	}

	var result types.ApplicationInfo
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
