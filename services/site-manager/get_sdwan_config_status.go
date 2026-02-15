package sitemanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

func (c *Client) GetSDWANConfigStatus(ctx context.Context, id string) (*types.SingleResponse[json.RawMessage], error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sd-wan-configs/%s/status", id), nil)
	if err != nil {
		return nil, err
	}

	var result types.SingleResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
