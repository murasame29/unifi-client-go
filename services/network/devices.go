package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListAdoptedDevices(ctx context.Context, req types.ListAdoptedDevicesRequest) (*types.PaginatedResponse[types.AdoptedDevice], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.AdoptedDevice]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) AdoptDevice(ctx context.Context, req types.AdoptDeviceRequest) (*types.AdoptedDevice, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.AdoptedDevice
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ExecutePortAction(ctx context.Context, req types.ExecutePortActionRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/interfaces/ports/%d/actions", req.SiteID, req.DeviceID, req.PortIdx), nil, req)
	if err != nil {
		return err
	}

	return internal.Decode(resp, &struct{}{})
}

func (c *Client) ExecuteDeviceAction(ctx context.Context, req types.ExecuteDeviceActionRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/actions", req.SiteID, req.DeviceID), nil, req)
	if err != nil {
		return err
	}

	return internal.Decode(resp, &struct{}{})
}

func (c *Client) GetAdoptedDeviceDetails(ctx context.Context, req types.GetAdoptedDeviceDetailsRequest) (*types.AdoptedDevice, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s", req.SiteID, req.DeviceID), nil)
	if err != nil {
		return nil, err
	}

	var result types.AdoptedDevice
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) RemoveDevice(ctx context.Context, req types.RemoveDeviceRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s", req.SiteID, req.DeviceID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetLatestDeviceStatistics(ctx context.Context, req types.GetLatestDeviceStatisticsRequest) (*types.DeviceStatistics, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/devices/%s/statistics/latest", req.SiteID, req.DeviceID), nil)
	if err != nil {
		return nil, err
	}

	var result types.DeviceStatistics
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListPendingDevices(ctx context.Context, req types.ListPendingDevicesRequest) (*types.PaginatedResponse[types.PendingDevice], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, "/v1/pending-devices", query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.PendingDevice]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
