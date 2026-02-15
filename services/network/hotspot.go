package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListVouchers(ctx context.Context, req types.ListVouchersRequest) (*types.PaginatedResponse[types.Voucher], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Voucher]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GenerateVouchers(ctx context.Context, req types.GenerateVouchersRequest) (*types.GenerateVouchersResponse, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.GenerateVouchersResponse
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteVouchers(ctx context.Context, req types.DeleteVouchersRequest) (*types.DeleteVoucherResponse, error) {
	query := url.Values{}
	if req.Filter != "" {
		query.Set("filter", req.Filter)
	}

	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.DeleteVoucherResponse
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetVoucherDetails(ctx context.Context, req types.GetVoucherDetailsRequest) (*types.Voucher, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers/%s", req.SiteID, req.VoucherID), nil)
	if err != nil {
		return nil, err
	}

	var result types.Voucher
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteVoucher(ctx context.Context, req types.DeleteVoucherRequest) (*types.DeleteVoucherResponse, error) {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/hotspot/vouchers/%s", req.SiteID, req.VoucherID), nil)
	if err != nil {
		return nil, err
	}

	var result types.DeleteVoucherResponse
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
