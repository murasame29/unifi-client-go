package network

import (
	"context"
	"fmt"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) GetTrafficMatchingList(ctx context.Context, req types.GetTrafficMatchingListRequest) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", req.SiteID, req.ListID), nil)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTrafficMatchingList(ctx context.Context, req types.UpdateTrafficMatchingListRequest) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Put(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", req.SiteID, req.ListID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteTrafficMatchingList(ctx context.Context, req types.DeleteTrafficMatchingListRequest) error {
	resp, err := c.client.Delete(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists/%s", req.SiteID, req.ListID), nil)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()
	return nil
}

func (c *Client) ListTrafficMatchingLists(ctx context.Context, req types.ListTrafficMatchingListsRequest) (*types.PaginatedResponse[types.TrafficMatchingList], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists", req.SiteID), query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.TrafficMatchingList]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateTrafficMatchingList(ctx context.Context, req types.CreateTrafficMatchingListRequest) (*types.TrafficMatchingList, error) {
	resp, err := c.client.Post(ctx, fmt.Sprintf("/v1/sites/%s/traffic-matching-lists", req.SiteID), nil, req)
	if err != nil {
		return nil, err
	}

	var result types.TrafficMatchingList
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
