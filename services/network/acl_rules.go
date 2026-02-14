package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetACLRule(ctx context.Context, siteID, ruleID string) (*types.ACLRule, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/%s", siteID, ruleID), nil)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateACLRule(ctx context.Context, siteID, ruleID string, req json.RawMessage) (*types.ACLRule, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/%s", siteID, ruleID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteACLRule(ctx context.Context, siteID, ruleID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/%s", siteID, ruleID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) GetACLRuleOrdering(ctx context.Context, siteID string) (json.RawMessage, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/ordering", siteID), nil)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) UpdateACLRuleOrdering(ctx context.Context, siteID string, req json.RawMessage) (json.RawMessage, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules/ordering", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result json.RawMessage
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListACLRules(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.ACLRule], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.ACLRule]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateACLRule(ctx context.Context, siteID string, req json.RawMessage) (*types.ACLRule, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/acl/rules", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.ACLRule
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
