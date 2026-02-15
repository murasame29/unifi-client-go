package sitemanager

import (
	"context"
	"fmt"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

func (c *Client) GetHostByID(ctx context.Context, id string) (*types.SingleResponse[types.Host], error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/hosts/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var result types.SingleResponse[types.Host]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
