package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetFirewallZone(ctx context.Context, req types.GetFirewallZoneRequest) (*types.FirewallZone, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", req.SiteID, req.ZoneID), nil)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFirewallZone(ctx context.Context, req types.UpdateFirewallZoneRequest) (*types.FirewallZone, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", req.SiteID, req.ZoneID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFirewallZone(ctx context.Context, req types.DeleteFirewallZoneRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones/%s", req.SiteID, req.ZoneID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetFirewallPolicy(ctx context.Context, req types.GetFirewallPolicyRequest) (*types.FirewallPolicy, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", req.SiteID, req.PolicyID), nil)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFirewallPolicy(ctx context.Context, req types.UpdateFirewallPolicyRequest) (*types.FirewallPolicy, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", req.SiteID, req.PolicyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteFirewallPolicy(ctx context.Context, req types.DeleteFirewallPolicyRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", req.SiteID, req.PolicyID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) PatchFirewallPolicy(ctx context.Context, req types.PatchFirewallPolicyRequest) (*types.FirewallPolicy, error) {
	resp, err := c.client.Patch(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/%s", req.SiteID, req.PolicyID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetFirewallPolicyOrdering(ctx context.Context, req types.GetFirewallPolicyOrderingRequest) (*types.FirewallPolicyOrdering, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/ordering", req.SiteID), nil)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicyOrdering
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateFirewallPolicyOrdering(ctx context.Context, req types.UpdateFirewallPolicyOrderingRequest) (*types.FirewallPolicyOrdering, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies/ordering", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicyOrdering
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListFirewallZones(ctx context.Context, req types.ListFirewallZonesRequest) (*types.PaginatedResponse[types.FirewallZone], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.FirewallZone]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFirewallZone(ctx context.Context, req types.CreateFirewallZoneRequest) (*types.FirewallZone, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/firewall/zones", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallZone
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListFirewallPolicies(ctx context.Context, req types.ListFirewallPoliciesRequest) (*types.PaginatedResponse[types.FirewallPolicy], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.FirewallPolicy]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateFirewallPolicy(ctx context.Context, req types.CreateFirewallPolicyRequest) (*types.FirewallPolicy, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/firewall/policies", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.FirewallPolicy
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
