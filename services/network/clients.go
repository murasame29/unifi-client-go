package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ExecuteClientAction(ctx context.Context, req types.ExecuteClientActionRequest) (*types.ClientActionResponse, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/clients/%s/actions", req.SiteID, req.ClientID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ClientActionResponse
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListConnectedClients(ctx context.Context, req types.ListConnectedClientsRequest) (*types.PaginatedResponse[types.ConnectedClientOverview], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/clients", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.ConnectedClientOverview]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConnectedClientDetails(ctx context.Context, req types.GetConnectedClientDetailsRequest) (*types.ConnectedClientDetails, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/clients/%s", req.SiteID, req.ClientID), nil)
	if err != nil {
		return nil, err
	}

	var result types.ConnectedClientDetails
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
