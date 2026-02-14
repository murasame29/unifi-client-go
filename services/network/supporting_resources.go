package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListWANInterfaces(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[json.RawMessage], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wan-interfaces", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListVPNTunnels(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[json.RawMessage], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/vpn/tunnels", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListVPNServers(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[json.RawMessage], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/vpn/servers", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListRadiusProfiles(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[json.RawMessage], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/radius-profiles", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListDeviceTags(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[json.RawMessage], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/device-tags", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[json.RawMessage]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListDPICategories(ctx context.Context, siteID string) ([]json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dpi/categories", siteID), nil)
	if err != nil {
		return nil, err
	}

	var result []json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListDPIApplications(ctx context.Context, siteID string) ([]json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dpi/applications", siteID), nil)
	if err != nil {
		return nil, err
	}

	var result []json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListCountries(ctx context.Context, siteID string) ([]json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/countries", siteID), nil)
	if err != nil {
		return nil, err
	}

	var result []json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
