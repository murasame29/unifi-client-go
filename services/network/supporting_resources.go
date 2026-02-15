package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListWANInterfaces(ctx context.Context, req types.ListWANInterfacesRequest) (*types.PaginatedResponse[types.WANInterface], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/wan-interfaces", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.WANInterface]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListVPNTunnels(ctx context.Context, req types.ListVPNTunnelsRequest) (*types.PaginatedResponse[types.VPNTunnel], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/vpn/tunnels", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.VPNTunnel]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListVPNServers(ctx context.Context, req types.ListVPNServersRequest) (*types.PaginatedResponse[types.VPNServer], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/vpn/servers", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.VPNServer]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListRadiusProfiles(ctx context.Context, req types.ListRadiusProfilesRequest) (*types.PaginatedResponse[types.RadiusProfile], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/radius-profiles", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.RadiusProfile]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListDeviceTags(ctx context.Context, req types.ListDeviceTagsRequest) (*types.PaginatedResponse[types.DeviceTag], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/device-tags", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.DeviceTag]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListDPICategories(ctx context.Context, req types.ListDPICategoriesRequest) ([]types.DPIApplicationCategory, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dpi/categories", req.SiteID), nil)
	if err != nil {
		return nil, err
	}

	var result []types.DPIApplicationCategory
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListDPIApplications(ctx context.Context, req types.ListDPIApplicationsRequest) ([]types.DPIApplication, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dpi/applications", req.SiteID), nil)
	if err != nil {
		return nil, err
	}

	var result []types.DPIApplication
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListCountries(ctx context.Context, req types.ListCountriesRequest) ([]types.Country, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/countries", req.SiteID), nil)
	if err != nil {
		return nil, err
	}

	var result []types.Country
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
