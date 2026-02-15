package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetWifiBroadcastDetails(ctx context.Context, req types.GetWifiBroadcastDetailsRequest) (*types.WifiBroadcast, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wifi/broadcasts/%s", req.SiteID, req.WifiBroadcastID), nil)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateWifiBroadcast(ctx context.Context, req types.UpdateWifiBroadcastRequest) (*types.WifiBroadcast, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/wifi/broadcasts/%s", req.SiteID, req.WifiBroadcastID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteWifiBroadcast(ctx context.Context, req types.DeleteWifiBroadcastRequest) error {
	query := url.Values{}
	if req.Force != nil && *req.Force {
		query.Set("force", "true")
	}

	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/wifi/broadcasts/%s", req.SiteID, req.WifiBroadcastID), query)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListWifiBroadcasts(ctx context.Context, req types.ListWifiBroadcastsRequest) (*types.PaginatedResponse[types.WifiBroadcast], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wifi/broadcasts", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.WifiBroadcast]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateWifiBroadcast(ctx context.Context, req types.CreateWifiBroadcastRequest) (*types.WifiBroadcast, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/wifi/broadcasts", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.WifiBroadcast
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
