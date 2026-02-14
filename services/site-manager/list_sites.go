package sitemanager

import (
	"context"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

type ListSitesParams struct {
	PageSize  string
	NextToken string
}

func (c *Client) ListSites(ctx context.Context, params *ListSitesParams) (*types.PaginatedResponse[types.Site], error) {
	query := url.Values{}
	if params != nil {
		if params.PageSize != "" {
			query.Set("pageSize", params.PageSize)
		}
		if params.NextToken != "" {
			query.Set("nextToken", params.NextToken)
		}
	}

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
