package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListAdoptedDevices(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.AdoptedDevice], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.AdoptedDevice]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AdoptDevice(ctx context.Context, siteID string, req types.AdoptDeviceRequest) (*types.AdoptedDevice, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.AdoptedDevice
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ExecutePortAction(ctx context.Context, siteID, deviceID string, portIdx int, req types.PortActionRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/ports/%d/actions", siteID, deviceID, portIdx), nil, req)
	if err != nil {
		return err
	}

	return internal.Decode(resp, &struct{}{})
}

func (c *Client) ExecuteDeviceAction(ctx context.Context, siteID, deviceID string, req types.DeviceActionRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/actions", siteID, deviceID), nil, req)
	if err != nil {
		return err
	}

	return internal.Decode(resp, &struct{}{})
}

func (c *Client) GetAdoptedDeviceDetails(ctx context.Context, siteID, deviceID string) (*types.AdoptedDevice, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s", siteID, deviceID), nil)
	if err != nil {
		return nil, err
	}

	var result types.AdoptedDevice
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveDevice(ctx context.Context, siteID, deviceID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s", siteID, deviceID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetLatestDeviceStatistics(ctx context.Context, siteID, deviceID string) (*types.DeviceStatistics, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/statistics/latest", siteID, deviceID), nil)
	if err != nil {
		return nil, err
	}

	var result types.DeviceStatistics
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListPendingDevices(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.PendingDevice], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices/pending", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.PendingDevice]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
