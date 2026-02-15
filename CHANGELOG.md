# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [services/site-manager/v1.0.0] - 2026-02-16

### Added

- Initial release of Site Manager API client
- Implemented endpoints:
  - `ListSites` - List all sites
  - `ListHosts` - List all hosts
  - `GetHostByID` - Get host details by ID
  - `ListDevices` - List devices with filtering
  - `GetISPMetrics` - Get ISP metrics by type
  - `QueryISPMetrics` - Query ISP metrics with filters
  - `ListSDWANConfigs` - List SD-WAN configurations
  - `GetSDWANConfigByID` - Get SD-WAN config by ID
  - `GetSDWANConfigStatus` - Get SD-WAN config status

## [services/network/v10.1.84] - 2026-02-16

### Added

- Initial release of Network API client (based on API v10.1.84)
- Implemented 66 endpoints across categories:
  - Application Info
  - Sites
  - Devices (adopted/pending, actions, statistics)
  - Clients (connected clients, actions, guest authorization)
  - Networks (CRUD, references)
  - WiFi Broadcasts (CRUD)
  - Hotspot (vouchers, operators)
  - Firewall (zones, policies, ordering)
  - ACL Rules (CRUD)
  - DNS Policies (CRUD)
  - Traffic Matching Lists (CRUD)
  - Supporting Resources (RADIUS profiles, device tags, port profiles)
