package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetDNSPolicy(ctx context.Context, req types.GetDNSPolicyRequest) (*types.DNSPolicy, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", req.SiteID, req.PolicyID), nil)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateDNSPolicy(ctx context.Context, req types.UpdateDNSPolicyRequest) (*types.DNSPolicy, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", req.SiteID, req.PolicyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteDNSPolicy(ctx context.Context, req types.DeleteDNSPolicyRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", req.SiteID, req.PolicyID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListDNSPolicies(ctx context.Context, req types.ListDNSPoliciesRequest) (*types.PaginatedResponse[types.DNSPolicy], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.DNSPolicy]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateDNSPolicy(ctx context.Context, req types.CreateDNSPolicyRequest) (*types.DNSPolicy, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
