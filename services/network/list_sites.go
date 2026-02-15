package network

import (
	"context"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func (c *Client) ListSites(ctx context.Context, req types.ListSitesRequest) (*types.PaginatedResponse[types.Site], error) {
	query := url.Values{}
	applyPagination(query, req.Pagination)

	resp, err := c.client.Get(ctx, "/v1/sites", query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Site]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
