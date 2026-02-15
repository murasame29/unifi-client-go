package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetACLRule(ctx context.Context, req types.GetACLRuleRequest) (*types.ACLRule, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl-rules/%s", req.SiteID, req.RuleID), nil)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateACLRule(ctx context.Context, req types.UpdateACLRuleRequest) (*types.ACLRule, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/acl-rules/%s", req.SiteID, req.RuleID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteACLRule(ctx context.Context, req types.DeleteACLRuleRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/acl-rules/%s", req.SiteID, req.RuleID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetACLRuleOrdering(ctx context.Context, req types.GetACLRuleOrderingRequest) (*types.ACLRuleOrdering, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/ordering", req.SiteID), nil)
	if err != nil {
		return nil, err
	}

	var result types.ACLRuleOrdering
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateACLRuleOrdering(ctx context.Context, req types.UpdateACLRuleOrderingRequest) (*types.ACLRuleOrdering, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/ordering", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ACLRuleOrdering
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListACLRules(ctx context.Context, req types.ListACLRulesRequest) (*types.PaginatedResponse[types.ACLRule], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.ACLRule]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateACLRule(ctx context.Context, req types.CreateACLRuleRequest) (*types.ACLRule, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
