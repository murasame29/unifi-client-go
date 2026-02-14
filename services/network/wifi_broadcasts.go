package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetWifiBroadcastDetails(ctx context.Context, siteID, wifiBroadcastID string) (*types.WifiBroadcast, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wifi-broadcasts/%s", siteID, wifiBroadcastID), nil)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateWifiBroadcast(ctx context.Context, siteID, wifiBroadcastID string, req types.UpdateWifiBroadcastRequest) (*types.WifiBroadcast, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/wifi-broadcasts/%s", siteID, wifiBroadcastID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteWifiBroadcast(ctx context.Context, siteID, wifiBroadcastID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/wifi-broadcasts/%s", siteID, wifiBroadcastID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListWifiBroadcasts(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.WifiBroadcast], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wifi-broadcasts", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.WifiBroadcast]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateWifiBroadcast(ctx context.Context, siteID string, req types.CreateWifiBroadcastRequest) (*types.WifiBroadcast, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/wifi-broadcasts", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
