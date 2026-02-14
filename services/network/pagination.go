package network

import (
	"net/url"
	"strconv"

	"github.com/murasame29/unifi-client-go/services/network/types"
)

func applyPagination(query url.Values, params *types.PaginationParams) {
	if params == nil {
		return
	}
	if params.Offset > 0 {
		query.Set("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit > 0 {
		query.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Filter != "" {
		query.Set("filter", params.Filter)
	}
}
