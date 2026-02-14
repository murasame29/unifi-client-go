package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetDNSPolicy(ctx context.Context, siteID, policyID string) (*types.DNSPolicy, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", siteID, policyID), nil)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateDNSPolicy(ctx context.Context, siteID, policyID string, req json.RawMessage) (*types.DNSPolicy, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", siteID, policyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteDNSPolicy(ctx context.Context, siteID, policyID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies/%s", siteID, policyID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListDNSPolicies(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.DNSPolicy], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.DNSPolicy]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateDNSPolicy(ctx context.Context, siteID string, req json.RawMessage) (*types.DNSPolicy, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/dns/policies", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.DNSPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
