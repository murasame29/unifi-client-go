package sitemanager

import (
	"context"
	"net/url"

	"github.com/murasame29/unifi-client-go/internal"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

type ListHostsParams struct {
	PageSize  string
	NextToken string
}

func (c *Client) ListHosts(ctx context.Context, params *ListHostsParams) (*types.PaginatedResponse[types.Host], error) {
	query := url.Values{}
	if params != nil {
		if params.PageSize != "" {
			query.Set("pageSize", params.PageSize)
		}
		if params.NextToken != "" {
			query.Set("nextToken", params.NextToken)
		}
	}

	resp, err := c.client.Get(ctx, "/v1/hosts", query)
	if err != nil {
		return nil, err
	}

	var result types.PaginatedResponse[types.Host]
	if err := internal.Decode(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
