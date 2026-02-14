package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetFirewallZone(ctx context.Context, siteID, zoneID string) (*types.FirewallZone, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", siteID, zoneID), nil)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFirewallZone(ctx context.Context, siteID, zoneID string, req types.UpdateFirewallZoneRequest) (*types.FirewallZone, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", siteID, zoneID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFirewallZone(ctx context.Context, siteID, zoneID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", siteID, zoneID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetFirewallPolicy(ctx context.Context, siteID, policyID string) (*types.FirewallPolicy, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", siteID, policyID), nil)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFirewallPolicy(ctx context.Context, siteID, policyID string, req json.RawMessage) (*types.FirewallPolicy, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", siteID, policyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFirewallPolicy(ctx context.Context, siteID, policyID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", siteID, policyID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) PatchFirewallPolicy(ctx context.Context, siteID, policyID string, req json.RawMessage) (*types.FirewallPolicy, error) {
	resp, err := c.client.Patch(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", siteID, policyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFirewallPolicyOrdering(ctx context.Context, siteID string) (json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/ordering", siteID), nil)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) UpdateFirewallPolicyOrdering(ctx context.Context, siteID string, req json.RawMessage) (json.RawMessage, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/ordering", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListFirewallZones(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.FirewallZone], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.FirewallZone]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFirewallZone(ctx context.Context, siteID string, req types.CreateFirewallZoneRequest) (*types.FirewallZone, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListFirewallPolicies(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.FirewallPolicy], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.FirewallPolicy]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFirewallPolicy(ctx context.Context, siteID string, req json.RawMessage) (*types.FirewallPolicy, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
