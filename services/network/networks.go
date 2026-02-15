package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetNetworkDetails(ctx context.Context, req types.GetNetworkDetailsRequest) (*types.Network, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", req.SiteID, req.NetworkID), nil)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateNetwork(ctx context.Context, req types.UpdateNetworkRequest) (*types.Network, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", req.SiteID, req.NetworkID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteNetwork(ctx context.Context, req types.DeleteNetworkRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s", req.SiteID, req.NetworkID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListNetworks(ctx context.Context, req types.ListNetworksRequest) (*types.PaginatedResponse[types.Network], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Network]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateNetwork(ctx context.Context, req types.CreateNetworkRequest) (*types.Network, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/networks", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.Network
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetNetworkReferences(ctx context.Context, req types.GetNetworkReferencesRequest) (json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/networks/%s/references", req.SiteID, req.NetworkID), nil)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
