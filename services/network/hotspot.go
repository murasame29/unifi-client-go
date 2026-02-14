package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListVouchers(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.Voucher], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Voucher]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GenerateVouchers(ctx context.Context, siteID string, req types.GenerateVouchersRequest) ([]types.Voucher, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result []types.Voucher
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) DeleteVouchers(ctx context.Context, siteID string, req types.DeleteVouchersRequest) error {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers/delete", siteID), nil, req)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetVoucherDetails(ctx context.Context, siteID, voucherID string) (*types.Voucher, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers/%s", siteID, voucherID), nil)
	if err != nil {
		return nil, err
	}

	var result types.Voucher
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteVoucher(ctx context.Context, siteID, voucherID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers/%s", siteID, voucherID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}
