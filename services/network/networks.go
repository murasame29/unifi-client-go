package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetNetworkDetails(ctx context.Context, siteID, networkID string) (*types.Network, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", siteID, networkID), nil)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateNetwork(ctx context.Context, siteID, networkID string, req types.UpdateNetworkRequest) (*types.Network, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", siteID, networkID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteNetwork(ctx context.Context, siteID, networkID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", siteID, networkID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListNetworks(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.Network], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Network]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateNetwork(ctx context.Context, siteID string, req types.CreateNetworkRequest) (*types.Network, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/networks", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetNetworkReferences(ctx context.Context, siteID, networkID string) (json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s/references", siteID, networkID), nil)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
