package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ExecuteClientAction(ctx context.Context, siteID, clientID string, req types.ClientActionRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/clients/%s/actions", siteID, clientID), nil, req)
	if err != nil {
		return err
	}

	return internal.Decode(resp, &struct{}{})
}

func (c *Client) ListConnectedClients(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.ConnectedClient], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/clients", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.ConnectedClient]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetConnectedClientDetails(ctx context.Context, siteID, clientID string) (*types.ConnectedClient, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/clients/%s", siteID, clientID), nil)
	if err != nil {
		return nil, err
	}

	var result types.ConnectedClient
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
