# unifi-client-go

Unofficial Go client library for UniFi API.

This library is based on the official [UniFi API documentation](https://developer.ui.com/).

## Supported APIs

| Service | API Version | Package |
|---------|-------------|---------|
| Site Manager API | v1 | `services/site-manager` |
| Network API | v10.1.84 | `services/network` |

## Installation

```bash
go get github.com/murasame29/unifi-client-go
```

## Usage

### Site Manager API

```go
package main

import (
	"context"
	"fmt"
	"log"

	sitemanager "github.com/murasame29/unifi-client-go/services/site-manager"
	"github.com/murasame29/unifi-client-go/services/site-manager/types"
)

func main() {
	client := sitemanager.NewClient("your-api-key")

	ctx := context.Background()

	// List sites
	sites, err := client.ListSites(ctx, types.ListSitesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, site := range sites.Data {
		fmt.Printf("Site: %s (%s)\n", site.Name, site.HostID)
	}
}
```

### Network API

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/murasame29/unifi-client-go/services/network"
	"github.com/murasame29/unifi-client-go/services/network/types"
)

func main() {
	client := network.NewClient("your-api-key")

	ctx := context.Background()

	// List sites
	sites, err := client.ListSites(ctx, types.ListSitesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, site := range sites.Data {
		fmt.Printf("Site: %s\n", site.Name)

		// List networks
		networks, err := client.ListNetworks(ctx, types.ListNetworksRequest{
			SiteID: site.ID,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, n := range networks.Data {
			fmt.Printf("  Network: %s (VLAN %d)\n", n.Name, n.VlanID)
		}
	}
}
```

### Options

You can specify a custom HTTP client or base URL:

```go
import "net/http"

// Custom HTTP client
httpClient := &http.Client{
	Timeout: 30 * time.Second,
}

client := network.NewClient(
	"your-api-key",
	network.WithHTTPClient(httpClient),
	network.WithBaseURL("https://custom-api.example.com"),
)
```

## Getting an API Key

Get your API key from UniFi Site Manager:
https://unifi.ui.com/


