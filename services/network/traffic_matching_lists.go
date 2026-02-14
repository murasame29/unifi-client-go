package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetTrafficMatchingList(ctx context.Context, siteID, listID string) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", siteID, listID), nil)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTrafficMatchingList(ctx context.Context, siteID, listID string, req json.RawMessage) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", siteID, listID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteTrafficMatchingList(ctx context.Context, siteID, listID string) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", siteID, listID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListTrafficMatchingLists(ctx context.Context, siteID string, params *types.PaginationParams) (*types.PaginatedResponse[types.TrafficMatchingList], error) {
	query := url.Values{}
	applyPagination(query, params)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists", siteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.TrafficMatchingList]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateTrafficMatchingList(ctx context.Context, siteID string, req json.RawMessage) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists", siteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
