package types

import "encoding/json"

type PaginatedResponse[T any] struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Count      int `json:"count"`
	TotalCount int `json:"totalCount"`
	Data       []T `json:"data"`
}

type PaginationParams struct {
	Offset int
	Limit  int
	Filter string
}

type EntityMetadata struct {
	Origin string `json:"origin"`
}

type ApplicationInfo struct {
	ApplicationVersion string `json:"applicationVersion"`
}

type Site struct {
	ID                string `json:"id"`
	InternalReference string `json:"internalReference"`
	Name              string `json:"name"`
}

type ListSitesRequest struct {
	Pagination *PaginationParams `json:"-"`
}

type DeviceUplink struct {
	DeviceID string `json:"deviceId"`
}

type DeviceFeatures struct {
	Switching   json.RawMessage `json:"switching"`
	AccessPoint json.RawMessage `json:"accessPoint"`
}

type PortPoE struct {
	Standard string `json:"standard"`
	Type     int    `json:"type"`
	Enabled  bool   `json:"enabled"`
	State    string `json:"state"`
}

type DevicePort struct {
	Idx          int      `json:"idx"`
	State        string   `json:"state"`
	Connector    string   `json:"connector"`
	MaxSpeedMbps int      `json:"maxSpeedMbps"`
	SpeedMbps    int      `json:"speedMbps,omitempty"`
	PoE          *PortPoE `json:"poe,omitempty"`
}

type DeviceRadio struct {
	WlanStandard    string  `json:"wlanStandard"`
	FrequencyGHz    float64 `json:"frequencyGHz"`
	ChannelWidthMHz int     `json:"channelWidthMHz"`
	Channel         int     `json:"channel,omitempty"`
}

type DeviceInterfaces struct {
	Ports  []DevicePort  `json:"ports,omitempty"`
	Radios []DeviceRadio `json:"radios,omitempty"`
}

type AdoptedDevice struct {
	ID                string            `json:"id"`
	MacAddress        string            `json:"macAddress"`
	IPAddress         string            `json:"ipAddress"`
	Name              string            `json:"name"`
	Model             string            `json:"model"`
	Supported         bool              `json:"supported"`
	State             string            `json:"state"`
	FirmwareVersion   string            `json:"firmwareVersion,omitempty"`
	FirmwareUpdatable bool              `json:"firmwareUpdatable"`
	AdoptedAt         string            `json:"adoptedAt,omitempty"`
	ProvisionedAt     string            `json:"provisionedAt,omitempty"`
	ConfigurationID   string            `json:"configurationId"`
	Uplink            *DeviceUplink     `json:"uplink,omitempty"`
	Features          *DeviceFeatures   `json:"features"`
	Interfaces        *DeviceInterfaces `json:"interfaces"`
}

type ListAdoptedDevicesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type AdoptDeviceRequest struct {
	SiteID            string `json:"-"`
	MacAddress        string `json:"macAddress"`
	IgnoreDeviceLimit bool   `json:"ignoreDeviceLimit"`
}

type ExecutePortActionRequest struct {
	SiteID   string `json:"-"`
	DeviceID string `json:"-"`
	PortIdx  int    `json:"-"`
	Action   string `json:"action"`
}

type ExecuteDeviceActionRequest struct {
	SiteID   string `json:"-"`
	DeviceID string `json:"-"`
	Action   string `json:"action"`
}

type GetAdoptedDeviceDetailsRequest struct {
	SiteID   string `json:"-"`
	DeviceID string `json:"-"`
}

type RemoveDeviceRequest struct {
	SiteID   string `json:"-"`
	DeviceID string `json:"-"`
}

type GetLatestDeviceStatisticsRequest struct {
	SiteID   string `json:"-"`
	DeviceID string `json:"-"`
}

type ListPendingDevicesRequest struct {
	Pagination *PaginationParams `json:"-"`
}

type StatisticsUplink struct {
	TxRateBps int `json:"txRateBps"`
	RxRateBps int `json:"rxRateBps"`
}

type StatisticsRadio struct {
	FrequencyGHz float64 `json:"frequencyGHz"`
	TxRetriesPct float64 `json:"txRetriesPct"`
}

type StatisticsInterfaces struct {
	Radios []StatisticsRadio `json:"radios,omitempty"`
}

type DeviceStatistics struct {
	UptimeSec            int                   `json:"uptimeSec"`
	LastHeartbeatAt      string                `json:"lastHeartbeatAt"`
	NextHeartbeatAt      string                `json:"nextHeartbeatAt"`
	LoadAverage1Min      float64               `json:"loadAverage1Min"`
	LoadAverage5Min      float64               `json:"loadAverage5Min"`
	LoadAverage15Min     float64               `json:"loadAverage15Min"`
	CpuUtilizationPct    float64               `json:"cpuUtilizationPct"`
	MemoryUtilizationPct float64               `json:"memoryUtilizationPct"`
	Uplink               *StatisticsUplink     `json:"uplink,omitempty"`
	Interfaces           *StatisticsInterfaces `json:"interfaces"`
}

type PendingDevice struct {
	MacAddress            string   `json:"macAddress"`
	IPAddress             string   `json:"ipAddress"`
	Model                 string   `json:"model"`
	State                 string   `json:"state"`
	Supported             bool     `json:"supported"`
	FirmwareVersion       string   `json:"firmwareVersion,omitempty"`
	FirmwareUpdatable     bool     `json:"firmwareUpdatable"`
	Features              []string `json:"features,omitempty"`
	AdoptionTargetSiteIds []string `json:"adoptionTargetSiteIds,omitempty"`
}

type ClientAccess struct {
	Type          string              `json:"type"`
	Authorized    *bool               `json:"authorized,omitempty"`
	Authorization *GuestAuthorization `json:"authorization,omitempty"`
}

type ConnectedClientOverview struct {
	Type           string        `json:"type"`
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	ConnectedAt    string        `json:"connectedAt,omitempty"`
	IPAddress      string        `json:"ipAddress,omitempty"`
	Access         *ClientAccess `json:"access"`
	MacAddress     string        `json:"macAddress,omitempty"`
	UplinkDeviceID string        `json:"uplinkDeviceId,omitempty"`
}

type ConnectedClientDetails struct {
	Type           string        `json:"type"`
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	ConnectedAt    string        `json:"connectedAt,omitempty"`
	IPAddress      string        `json:"ipAddress,omitempty"`
	Access         *ClientAccess `json:"access"`
	MacAddress     string        `json:"macAddress,omitempty"`
	UplinkDeviceID string        `json:"uplinkDeviceId,omitempty"`
}

type ListConnectedClientsRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type GetConnectedClientDetailsRequest struct {
	SiteID   string `json:"-"`
	ClientID string `json:"-"`
}

type AuthorizationUsage struct {
	DurationSec int `json:"durationSec"`
	RxBytes     int `json:"rxBytes"`
	TxBytes     int `json:"txBytes"`
	Bytes       int `json:"bytes"`
}

type GuestAuthorization struct {
	AuthorizedAt         string              `json:"authorizedAt"`
	AuthorizationMethod  string              `json:"authorizationMethod"`
	ExpiresAt            string              `json:"expiresAt"`
	DataUsageLimitMBytes *int                `json:"dataUsageLimitMBytes,omitempty"`
	RxRateLimitKbps      *int                `json:"rxRateLimitKbps,omitempty"`
	TxRateLimitKbps      *int                `json:"txRateLimitKbps,omitempty"`
	Usage                *AuthorizationUsage `json:"usage,omitempty"`
}

type ClientActionResponse struct {
	Action               string              `json:"action"`
	RevokedAuthorization *GuestAuthorization `json:"revokedAuthorization,omitempty"`
	GrantedAuthorization *GuestAuthorization `json:"grantedAuthorization,omitempty"`
}

type ExecuteClientActionRequest struct {
	SiteID               string `json:"-"`
	ClientID             string `json:"-"`
	Action               string `json:"action"`
	TimeLimitMinutes     *int   `json:"timeLimitMinutes,omitempty"`
	DataUsageLimitMBytes *int   `json:"dataUsageLimitMBytes,omitempty"`
	RxRateLimitKbps      *int   `json:"rxRateLimitKbps,omitempty"`
	TxRateLimitKbps      *int   `json:"txRateLimitKbps,omitempty"`
}

type DhcpGuarding struct {
	TrustedDhcpServerIPAddresses []string `json:"trustedDhcpServerIpAddresses,omitempty"`
}

type NetworkDHCPIPAddressRange struct {
	Start string `json:"start,omitempty"`
	Stop  string `json:"stop,omitempty"`
}

type NetworkPXEConfiguration struct {
	ServerIPAddress string `json:"serverIpAddress"`
	Filename        string `json:"filename"`
}

type NetworkDHCPConfiguration struct {
	Mode                         string                     `json:"mode"`
	IPAddressRange               *NetworkDHCPIPAddressRange `json:"ipAddressRange,omitempty"`
	GatewayIPAddressOverride     string                     `json:"gatewayIpAddressOverride,omitempty"`
	DnsServerIPAddressesOverride []string                   `json:"dnsServerIpAddressesOverride,omitempty"`
	LeaseTimeSeconds             *int                       `json:"leaseTimeSeconds,omitempty"`
	DomainName                   string                     `json:"domainName,omitempty"`
	PingConflictDetectionEnabled *bool                      `json:"pingConflictDetectionEnabled,omitempty"`
	PxeConfiguration             *NetworkPXEConfiguration   `json:"pxeConfiguration,omitempty"`
	NtpServerIPAddresses         []string                   `json:"ntpServerIpAddresses,omitempty"`
	Option43Value                string                     `json:"option43Value,omitempty"`
	TftpServerAddress            string                     `json:"tftpServerAddress,omitempty"`
	TimeOffsetSeconds            *int                       `json:"timeOffsetSeconds,omitempty"`
	WpadUrl                      string                     `json:"wpadUrl,omitempty"`
	WinsServerIPAddresses        []string                   `json:"winsServerIpAddresses,omitempty"`
	DhcpServerIPAddresses        []string                   `json:"dhcpServerIpAddresses,omitempty"`
}

type IPAddressSelector struct {
	Type  string `json:"type"`
	Value string `json:"value,omitempty"`
}

type NetworkNATOutboundIPAddressConfig struct {
	Type               string              `json:"type"`
	WanInterfaceID     string              `json:"wanInterfaceId"`
	IpAddressSelectors []IPAddressSelector `json:"ipAddressSelectors,omitempty"`
}

type NetworkIPv4Configuration struct {
	AutoScaleEnabled                  *bool                               `json:"autoScaleEnabled,omitempty"`
	HostIPAddress                     string                              `json:"hostIpAddress,omitempty"`
	PrefixLength                      *int                                `json:"prefixLength,omitempty"`
	AdditionalHostIPSubnets           []string                            `json:"additionalHostIpSubnets,omitempty"`
	DhcpConfiguration                 *NetworkDHCPConfiguration           `json:"dhcpConfiguration,omitempty"`
	NatOutboundIPAddressConfiguration []NetworkNATOutboundIPAddressConfig `json:"natOutboundIpAddressConfiguration,omitempty"`
}

type IPv6AddressSuffixRange struct {
	Start string `json:"start,omitempty"`
	Stop  string `json:"stop,omitempty"`
}

type IPv6DHCPConfiguration struct {
	IPAddressSuffixRange *IPv6AddressSuffixRange `json:"ipAddressSuffixRange"`
	LeaseTimeSeconds     int                     `json:"leaseTimeSeconds"`
}

type IPv6ClientAddressAssignment struct {
	DhcpConfiguration *IPv6DHCPConfiguration `json:"dhcpConfiguration,omitempty"`
	SlaacEnabled      bool                   `json:"slaacEnabled"`
}

type IPv6RouterAdvertisement struct {
	Priority string `json:"priority"`
}

type NetworkIPv6Configuration struct {
	InterfaceType                  string                       `json:"interfaceType"`
	ClientAddressAssignment        *IPv6ClientAddressAssignment `json:"clientAddressAssignment"`
	RouterAdvertisement            *IPv6RouterAdvertisement     `json:"routerAdvertisement,omitempty"`
	DnsServerIPAddressesOverride   []string                     `json:"dnsServerIpAddressesOverride,omitempty"`
	AdditionalHostIPSubnets        []string                     `json:"additionalHostIpSubnets,omitempty"`
	PrefixDelegationWanInterfaceID string                       `json:"prefixDelegationWanInterfaceId,omitempty"`
}

type Network struct {
	Management            string                    `json:"management"`
	ID                    string                    `json:"id,omitempty"`
	Name                  string                    `json:"name"`
	Enabled               bool                      `json:"enabled"`
	VlanID                int                       `json:"vlanId"`
	Metadata              *EntityMetadata           `json:"metadata,omitempty"`
	DhcpGuarding          *DhcpGuarding             `json:"dhcpGuarding,omitempty"`
	Default               bool                      `json:"default,omitempty"`
	IsolationEnabled      *bool                     `json:"isolationEnabled,omitempty"`
	CellularBackupEnabled *bool                     `json:"cellularBackupEnabled,omitempty"`
	DeviceID              string                    `json:"deviceId,omitempty"`
	ZoneID                string                    `json:"zoneId,omitempty"`
	InternetAccessEnabled *bool                     `json:"internetAccessEnabled,omitempty"`
	MdnsForwardingEnabled *bool                     `json:"mdnsForwardingEnabled,omitempty"`
	IPv4Configuration     *NetworkIPv4Configuration `json:"ipv4Configuration,omitempty"`
	IPv6Configuration     *NetworkIPv6Configuration `json:"ipv6Configuration,omitempty"`
}

type ListNetworksRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateNetworkRequest struct {
	SiteID                string                    `json:"-"`
	Management            string                    `json:"management"`
	Name                  string                    `json:"name"`
	Enabled               bool                      `json:"enabled"`
	VlanID                int                       `json:"vlanId"`
	DhcpGuarding          *DhcpGuarding             `json:"dhcpGuarding,omitempty"`
	IsolationEnabled      *bool                     `json:"isolationEnabled,omitempty"`
	CellularBackupEnabled *bool                     `json:"cellularBackupEnabled,omitempty"`
	DeviceID              string                    `json:"deviceId,omitempty"`
	ZoneID                string                    `json:"zoneId,omitempty"`
	InternetAccessEnabled *bool                     `json:"internetAccessEnabled,omitempty"`
	MdnsForwardingEnabled *bool                     `json:"mdnsForwardingEnabled,omitempty"`
	IPv4Configuration     *NetworkIPv4Configuration `json:"ipv4Configuration,omitempty"`
	IPv6Configuration     *NetworkIPv6Configuration `json:"ipv6Configuration,omitempty"`
}

type GetNetworkDetailsRequest struct {
	SiteID    string `json:"-"`
	NetworkID string `json:"-"`
}

type UpdateNetworkRequest struct {
	SiteID                string                    `json:"-"`
	NetworkID             string                    `json:"-"`
	Management            string                    `json:"management"`
	Name                  string                    `json:"name"`
	Enabled               bool                      `json:"enabled"`
	VlanID                int                       `json:"vlanId"`
	DhcpGuarding          *DhcpGuarding             `json:"dhcpGuarding,omitempty"`
	IsolationEnabled      *bool                     `json:"isolationEnabled,omitempty"`
	CellularBackupEnabled *bool                     `json:"cellularBackupEnabled,omitempty"`
	DeviceID              string                    `json:"deviceId,omitempty"`
	ZoneID                string                    `json:"zoneId,omitempty"`
	InternetAccessEnabled *bool                     `json:"internetAccessEnabled,omitempty"`
	MdnsForwardingEnabled *bool                     `json:"mdnsForwardingEnabled,omitempty"`
	IPv4Configuration     *NetworkIPv4Configuration `json:"ipv4Configuration,omitempty"`
	IPv6Configuration     *NetworkIPv6Configuration `json:"ipv6Configuration,omitempty"`
}

type DeleteNetworkRequest struct {
	SiteID    string `json:"-"`
	NetworkID string `json:"-"`
	Force     *bool  `json:"-"`
}

type GetNetworkReferencesRequest struct {
	SiteID    string `json:"-"`
	NetworkID string `json:"-"`
}

type NetworkReferenceDetail struct {
	ReferenceID string `json:"referenceId"`
}

type NetworkReferenceResource struct {
	ResourceType   string                   `json:"resourceType"`
	ReferenceCount int                      `json:"referenceCount"`
	References     []NetworkReferenceDetail `json:"references,omitempty"`
}

type NetworkReferencesResponse struct {
	ReferenceResources []NetworkReferenceResource `json:"referenceResources"`
}

type WifiNetworkReference struct {
	Type      string `json:"type,omitempty"`
	NetworkID string `json:"networkId,omitempty"`
}

type WifiSecurityConfiguration struct {
	Type                      string                   `json:"type"`
	RadiusConfiguration       *WifiRadiusConfiguration `json:"radiusConfiguration,omitempty"`
	Passphrase                string                   `json:"passphrase,omitempty"`
	GroupRekeyIntervalSeconds *int                     `json:"groupRekeyIntervalSeconds,omitempty"`
	FastRoamingEnabled        *bool                    `json:"fastRoamingEnabled,omitempty"`
	PmfMode                   string                   `json:"pmfMode,omitempty"`
	PresharedKeys             []WifiPresharedKey       `json:"presharedKeys,omitempty"`
	SaeConfiguration          *SaeConfiguration        `json:"saeConfiguration,omitempty"`
	CoaEnabled                *bool                    `json:"coaEnabled,omitempty"`
	SecurityMode              string                   `json:"securityMode,omitempty"`
	Wpa3FastRoamingEnabled    *bool                    `json:"wpa3FastRoamingEnabled,omitempty"`
}

type WifiPresharedKey struct {
	Name       string `json:"name,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
}

type SaeConfiguration struct {
	AnticloggingThresholdSeconds int `json:"anticloggingThresholdSeconds"`
	SyncTimeSeconds              int `json:"syncTimeSeconds"`
}

type WifiRadiusConfiguration struct {
	ProfileID                      string                          `json:"profileId"`
	NasId                          *WifiRadiusNasIdConfiguration   `json:"nasId,omitempty"`
	MacAuthenticationConfiguration *WifiRadiusMacAuthConfiguration `json:"macAuthenticationConfiguration,omitempty"`
}

type WifiRadiusNasIdConfiguration struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type WifiRadiusMacAuthConfiguration struct {
	Enabled          *bool  `json:"enabled,omitempty"`
	PasswordType     string `json:"passwordType,omitempty"`
	MacAddressFormat string `json:"macAddressFormat,omitempty"`
}

type BroadcastingDeviceFilter struct {
	Type         string   `json:"type"`
	DeviceIDs    []string `json:"deviceIds,omitempty"`
	DeviceTagIDs []string `json:"deviceTagIds,omitempty"`
}

type MdnsProxyConfiguration struct {
	Mode     string            `json:"mode"`
	Policies []MdnsProxyPolicy `json:"policies,omitempty"`
}

type MdnsProxyPolicy struct {
	Type      string   `json:"type,omitempty"`
	Services  []string `json:"services,omitempty"`
	NetworkID string   `json:"networkId,omitempty"`
}

type MulticastFilteringPolicy struct {
	Action                 string   `json:"action"`
	SourceMacAddressFilter []string `json:"sourceMacAddressFilter,omitempty"`
}

type BasicDataRateKbpsByFrequencyGHz struct {
	TwoPointFour int `json:"2.4,omitempty"`
	Five         int `json:"5,omitempty"`
}

type ClientFilteringPolicy struct {
	Action           string   `json:"action"`
	MacAddressFilter []string `json:"macAddressFilter,omitempty"`
}

type BlackoutScheduleDay struct {
	Type       string                      `json:"type"`
	Day        string                      `json:"day"`
	TimeRanges []BlackoutScheduleTimeRange `json:"timeRanges,omitempty"`
}

type BlackoutScheduleTimeRange struct {
	Start string `json:"start,omitempty"`
	Stop  string `json:"stop,omitempty"`
}

type BlackoutScheduleConfiguration struct {
	Days []BlackoutScheduleDay `json:"days,omitempty"`
}

type WifiHotspotConfiguration struct {
	Enabled bool   `json:"enabled,omitempty"`
	Type    string `json:"type,omitempty"`
}

type DtimPeriodByFrequencyGHzOverride struct {
	TwoPointFour *int `json:"2.4,omitempty"`
	Five         *int `json:"5,omitempty"`
	Six          *int `json:"6,omitempty"`
}

type WifiBroadcast struct {
	Type                                string                            `json:"type"`
	ID                                  string                            `json:"id,omitempty"`
	Name                                string                            `json:"name"`
	Metadata                            *EntityMetadata                   `json:"metadata,omitempty"`
	Enabled                             bool                              `json:"enabled"`
	Network                             *WifiNetworkReference             `json:"network,omitempty"`
	SecurityConfiguration               *WifiSecurityConfiguration        `json:"securityConfiguration,omitempty"`
	BroadcastingDeviceFilter            *BroadcastingDeviceFilter         `json:"broadcastingDeviceFilter,omitempty"`
	MdnsProxyConfiguration              *MdnsProxyConfiguration           `json:"mdnsProxyConfiguration,omitempty"`
	MulticastFilteringPolicy            *MulticastFilteringPolicy         `json:"multicastFilteringPolicy,omitempty"`
	MulticastToUnicastConversionEnabled bool                              `json:"multicastToUnicastConversionEnabled"`
	ClientIsolationEnabled              bool                              `json:"clientIsolationEnabled"`
	HideName                            bool                              `json:"hideName"`
	UapsdEnabled                        bool                              `json:"uapsdEnabled"`
	BasicDataRateKbpsByFrequencyGHz     *BasicDataRateKbpsByFrequencyGHz  `json:"basicDataRateKbpsByFrequencyGHz,omitempty"`
	ClientFilteringPolicy               *ClientFilteringPolicy            `json:"clientFilteringPolicy,omitempty"`
	BlackoutScheduleConfiguration       *BlackoutScheduleConfiguration    `json:"blackoutScheduleConfiguration,omitempty"`
	BroadcastingFrequenciesGHz          []float64                         `json:"broadcastingFrequenciesGHz,omitempty"`
	HotspotConfiguration                *WifiHotspotConfiguration         `json:"hotspotConfiguration,omitempty"`
	MloEnabled                          *bool                             `json:"mloEnabled,omitempty"`
	BandSteeringEnabled                 *bool                             `json:"bandSteeringEnabled,omitempty"`
	ArpProxyEnabled                     *bool                             `json:"arpProxyEnabled,omitempty"`
	BssTransitionEnabled                *bool                             `json:"bssTransitionEnabled,omitempty"`
	AdvertiseDeviceName                 *bool                             `json:"advertiseDeviceName,omitempty"`
	DtimPeriodByFrequencyGHzOverride    *DtimPeriodByFrequencyGHzOverride `json:"dtimPeriodByFrequencyGHzOverride,omitempty"`
}

type ListWifiBroadcastsRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateWifiBroadcastRequest struct {
	SiteID                              string                            `json:"-"`
	Type                                string                            `json:"type"`
	Name                                string                            `json:"name"`
	Enabled                             bool                              `json:"enabled"`
	Network                             *WifiNetworkReference             `json:"network,omitempty"`
	SecurityConfiguration               *WifiSecurityConfiguration        `json:"securityConfiguration,omitempty"`
	BroadcastingDeviceFilter            *BroadcastingDeviceFilter         `json:"broadcastingDeviceFilter,omitempty"`
	MdnsProxyConfiguration              *MdnsProxyConfiguration           `json:"mdnsProxyConfiguration,omitempty"`
	MulticastFilteringPolicy            *MulticastFilteringPolicy         `json:"multicastFilteringPolicy,omitempty"`
	MulticastToUnicastConversionEnabled bool                              `json:"multicastToUnicastConversionEnabled"`
	ClientIsolationEnabled              bool                              `json:"clientIsolationEnabled"`
	HideName                            bool                              `json:"hideName"`
	UapsdEnabled                        bool                              `json:"uapsdEnabled"`
	BasicDataRateKbpsByFrequencyGHz     *BasicDataRateKbpsByFrequencyGHz  `json:"basicDataRateKbpsByFrequencyGHz,omitempty"`
	ClientFilteringPolicy               *ClientFilteringPolicy            `json:"clientFilteringPolicy,omitempty"`
	BlackoutScheduleConfiguration       *BlackoutScheduleConfiguration    `json:"blackoutScheduleConfiguration,omitempty"`
	BroadcastingFrequenciesGHz          []float64                         `json:"broadcastingFrequenciesGHz,omitempty"`
	HotspotConfiguration                *WifiHotspotConfiguration         `json:"hotspotConfiguration,omitempty"`
	MloEnabled                          *bool                             `json:"mloEnabled,omitempty"`
	BandSteeringEnabled                 *bool                             `json:"bandSteeringEnabled,omitempty"`
	ArpProxyEnabled                     *bool                             `json:"arpProxyEnabled,omitempty"`
	BssTransitionEnabled                *bool                             `json:"bssTransitionEnabled,omitempty"`
	AdvertiseDeviceName                 *bool                             `json:"advertiseDeviceName,omitempty"`
	DtimPeriodByFrequencyGHzOverride    *DtimPeriodByFrequencyGHzOverride `json:"dtimPeriodByFrequencyGHzOverride,omitempty"`
}

type GetWifiBroadcastDetailsRequest struct {
	SiteID          string `json:"-"`
	WifiBroadcastID string `json:"-"`
}

type UpdateWifiBroadcastRequest struct {
	SiteID                              string                            `json:"-"`
	WifiBroadcastID                     string                            `json:"-"`
	Type                                string                            `json:"type"`
	Name                                string                            `json:"name"`
	Enabled                             bool                              `json:"enabled"`
	Network                             *WifiNetworkReference             `json:"network,omitempty"`
	SecurityConfiguration               *WifiSecurityConfiguration        `json:"securityConfiguration,omitempty"`
	BroadcastingDeviceFilter            *BroadcastingDeviceFilter         `json:"broadcastingDeviceFilter,omitempty"`
	MdnsProxyConfiguration              *MdnsProxyConfiguration           `json:"mdnsProxyConfiguration,omitempty"`
	MulticastFilteringPolicy            *MulticastFilteringPolicy         `json:"multicastFilteringPolicy,omitempty"`
	MulticastToUnicastConversionEnabled bool                              `json:"multicastToUnicastConversionEnabled"`
	ClientIsolationEnabled              bool                              `json:"clientIsolationEnabled"`
	HideName                            bool                              `json:"hideName"`
	UapsdEnabled                        bool                              `json:"uapsdEnabled"`
	BasicDataRateKbpsByFrequencyGHz     *BasicDataRateKbpsByFrequencyGHz  `json:"basicDataRateKbpsByFrequencyGHz,omitempty"`
	ClientFilteringPolicy               *ClientFilteringPolicy            `json:"clientFilteringPolicy,omitempty"`
	BlackoutScheduleConfiguration       *BlackoutScheduleConfiguration    `json:"blackoutScheduleConfiguration,omitempty"`
	BroadcastingFrequenciesGHz          []float64                         `json:"broadcastingFrequenciesGHz,omitempty"`
	HotspotConfiguration                *WifiHotspotConfiguration         `json:"hotspotConfiguration,omitempty"`
	MloEnabled                          *bool                             `json:"mloEnabled,omitempty"`
	BandSteeringEnabled                 *bool                             `json:"bandSteeringEnabled,omitempty"`
	ArpProxyEnabled                     *bool                             `json:"arpProxyEnabled,omitempty"`
	BssTransitionEnabled                *bool                             `json:"bssTransitionEnabled,omitempty"`
	AdvertiseDeviceName                 *bool                             `json:"advertiseDeviceName,omitempty"`
	DtimPeriodByFrequencyGHzOverride    *DtimPeriodByFrequencyGHzOverride `json:"dtimPeriodByFrequencyGHzOverride,omitempty"`
}

type DeleteWifiBroadcastRequest struct {
	SiteID          string `json:"-"`
	WifiBroadcastID string `json:"-"`
}

type Voucher struct {
	ID        string `json:"id,omitempty"`
	Code      string `json:"code,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	DataQuota int    `json:"dataQuota,omitempty"`
	Note      string `json:"note,omitempty"`
	Used      bool   `json:"used,omitempty"`
}

type ListVouchersRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type GenerateVouchersRequest struct {
	SiteID    string `json:"-"`
	Count     int    `json:"count"`
	Duration  int    `json:"duration"`
	DataQuota int    `json:"dataQuota,omitempty"`
	Note      string `json:"note,omitempty"`
}

type GetVoucherDetailsRequest struct {
	SiteID    string `json:"-"`
	VoucherID string `json:"-"`
}

type DeleteVoucherRequest struct {
	SiteID    string `json:"-"`
	VoucherID string `json:"-"`
}

type DeleteVouchersRequest struct {
	SiteID string   `json:"-"`
	IDs    []string `json:"ids"`
}

type FirewallZone struct {
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name"`
	NetworkIDs []string        `json:"networkIds,omitempty"`
	Metadata   *EntityMetadata `json:"metadata,omitempty"`
}

type ListFirewallZonesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateFirewallZoneRequest struct {
	SiteID     string   `json:"-"`
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds,omitempty"`
}

type GetFirewallZoneRequest struct {
	SiteID string `json:"-"`
	ZoneID string `json:"-"`
}

type UpdateFirewallZoneRequest struct {
	SiteID     string   `json:"-"`
	ZoneID     string   `json:"-"`
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds,omitempty"`
}

type DeleteFirewallZoneRequest struct {
	SiteID string `json:"-"`
	ZoneID string `json:"-"`
}

type FirewallPolicyAction struct {
	Type               string `json:"type"`
	AllowReturnTraffic *bool  `json:"allowReturnTraffic,omitempty"`
}

type TrafficFilter struct {
	Type             string                 `json:"type"`
	PortFilter       *FirewallPortFilter    `json:"portFilter,omitempty"`
	NetworkFilter    *FirewallNetworkFilter `json:"networkFilter,omitempty"`
	MacAddressFilter string                 `json:"macAddressFilter,omitempty"`
}

type FirewallPortFilter struct {
	Ports []int `json:"ports,omitempty"`
}

type FirewallNetworkFilter struct {
	NetworkIds []string `json:"networkIds,omitempty"`
}

type FirewallPolicyEndpoint struct {
	ZoneID        string         `json:"zoneId"`
	TrafficFilter *TrafficFilter `json:"trafficFilter,omitempty"`
}

type FirewallIPProtocolScope struct {
	IPVersion      string                  `json:"ipVersion"`
	ProtocolFilter *FirewallProtocolFilter `json:"protocolFilter,omitempty"`
}

type FirewallProtocolFilter struct {
	Type          string            `json:"type"`
	Protocol      *FirewallProtocol `json:"protocol,omitempty"`
	MatchOpposite *bool             `json:"matchOpposite,omitempty"`
}

type FirewallProtocol struct {
	Name string `json:"name,omitempty"`
}

type FirewallSchedule struct {
	Mode       string              `json:"mode"`
	TimeFilter *FirewallTimeFilter `json:"timeFilter,omitempty"`
}

type FirewallTimeFilter struct {
	StartTime string `json:"startTime,omitempty"`
	StopTime  string `json:"stopTime,omitempty"`
}

type FirewallPolicy struct {
	ID                    string                   `json:"id,omitempty"`
	Enabled               bool                     `json:"enabled"`
	Name                  string                   `json:"name"`
	Description           string                   `json:"description,omitempty"`
	Index                 int                      `json:"index,omitempty"`
	Action                *FirewallPolicyAction    `json:"action,omitempty"`
	Source                *FirewallPolicyEndpoint  `json:"source,omitempty"`
	Destination           *FirewallPolicyEndpoint  `json:"destination,omitempty"`
	IPProtocolScope       *FirewallIPProtocolScope `json:"ipProtocolScope,omitempty"`
	ConnectionStateFilter []string                 `json:"connectionStateFilter,omitempty"`
	IpsecFilter           string                   `json:"ipsecFilter,omitempty"`
	LoggingEnabled        bool                     `json:"loggingEnabled"`
	Schedule              *FirewallSchedule        `json:"schedule,omitempty"`
	Metadata              *EntityMetadata          `json:"metadata,omitempty"`
}

type ListFirewallPoliciesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateFirewallPolicyRequest struct {
	SiteID                string                   `json:"-"`
	Enabled               bool                     `json:"enabled"`
	Name                  string                   `json:"name"`
	Description           string                   `json:"description,omitempty"`
	Action                *FirewallPolicyAction    `json:"action"`
	Source                *FirewallPolicyEndpoint  `json:"source"`
	Destination           *FirewallPolicyEndpoint  `json:"destination"`
	IPProtocolScope       *FirewallIPProtocolScope `json:"ipProtocolScope"`
	ConnectionStateFilter []string                 `json:"connectionStateFilter,omitempty"`
	IpsecFilter           string                   `json:"ipsecFilter,omitempty"`
	LoggingEnabled        bool                     `json:"loggingEnabled"`
	Schedule              *FirewallSchedule        `json:"schedule,omitempty"`
}

type GetFirewallPolicyRequest struct {
	SiteID   string `json:"-"`
	PolicyID string `json:"-"`
}

type UpdateFirewallPolicyRequest struct {
	SiteID                string                   `json:"-"`
	PolicyID              string                   `json:"-"`
	Enabled               bool                     `json:"enabled"`
	Name                  string                   `json:"name"`
	Description           string                   `json:"description,omitempty"`
	Action                *FirewallPolicyAction    `json:"action"`
	Source                *FirewallPolicyEndpoint  `json:"source"`
	Destination           *FirewallPolicyEndpoint  `json:"destination"`
	IPProtocolScope       *FirewallIPProtocolScope `json:"ipProtocolScope"`
	ConnectionStateFilter []string                 `json:"connectionStateFilter,omitempty"`
	IpsecFilter           string                   `json:"ipsecFilter,omitempty"`
	LoggingEnabled        bool                     `json:"loggingEnabled"`
	Schedule              *FirewallSchedule        `json:"schedule,omitempty"`
}

type DeleteFirewallPolicyRequest struct {
	SiteID   string `json:"-"`
	PolicyID string `json:"-"`
}

type PatchFirewallPolicyRequest struct {
	SiteID                string                   `json:"-"`
	PolicyID              string                   `json:"-"`
	Enabled               *bool                    `json:"enabled,omitempty"`
	Name                  string                   `json:"name,omitempty"`
	Description           string                   `json:"description,omitempty"`
	Action                *FirewallPolicyAction    `json:"action,omitempty"`
	Source                *FirewallPolicyEndpoint  `json:"source,omitempty"`
	Destination           *FirewallPolicyEndpoint  `json:"destination,omitempty"`
	IPProtocolScope       *FirewallIPProtocolScope `json:"ipProtocolScope,omitempty"`
	ConnectionStateFilter []string                 `json:"connectionStateFilter,omitempty"`
	IpsecFilter           string                   `json:"ipsecFilter,omitempty"`
	LoggingEnabled        *bool                    `json:"loggingEnabled,omitempty"`
	Schedule              *FirewallSchedule        `json:"schedule,omitempty"`
}

type FirewallPolicyOrdering struct {
	PolicyIDs []string `json:"policyIds"`
}

type GetFirewallPolicyOrderingRequest struct {
	SiteID string `json:"-"`
}

type UpdateFirewallPolicyOrderingRequest struct {
	SiteID    string   `json:"-"`
	PolicyIDs []string `json:"policyIds"`
}

type ACLDeviceFilter struct {
	Type      string   `json:"type"`
	DeviceIDs []string `json:"deviceIds,omitempty"`
}

type ACLEndpointFilter struct {
	Type                 string   `json:"type"`
	IpAddressesOrSubnets []string `json:"ipAddressesOrSubnets,omitempty"`
	NetworkIds           []string `json:"networkIds,omitempty"`
	PortFilter           []int    `json:"portFilter,omitempty"`
	MacAddresses         []string `json:"macAddresses,omitempty"`
	PrefixLength         *int     `json:"prefixLength,omitempty"`
}

type ACLRule struct {
	Type                  string             `json:"type"`
	ID                    string             `json:"id,omitempty"`
	Enabled               bool               `json:"enabled"`
	Name                  string             `json:"name"`
	Description           string             `json:"description,omitempty"`
	Action                string             `json:"action"`
	EnforcingDeviceFilter *ACLDeviceFilter   `json:"enforcingDeviceFilter,omitempty"`
	Index                 int                `json:"index,omitempty"`
	SourceFilter          *ACLEndpointFilter `json:"sourceFilter,omitempty"`
	DestinationFilter     *ACLEndpointFilter `json:"destinationFilter,omitempty"`
	Metadata              *EntityMetadata    `json:"metadata,omitempty"`
	ProtocolFilter        []string           `json:"protocolFilter,omitempty"`
	NetworkIdFilter       string             `json:"networkIdFilter,omitempty"`
}

type ListACLRulesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateACLRuleRequest struct {
	SiteID                string             `json:"-"`
	Type                  string             `json:"type"`
	Enabled               bool               `json:"enabled"`
	Name                  string             `json:"name"`
	Description           string             `json:"description,omitempty"`
	Action                string             `json:"action"`
	EnforcingDeviceFilter *ACLDeviceFilter   `json:"enforcingDeviceFilter,omitempty"`
	Index                 int                `json:"index,omitempty"`
	SourceFilter          *ACLEndpointFilter `json:"sourceFilter,omitempty"`
	DestinationFilter     *ACLEndpointFilter `json:"destinationFilter,omitempty"`
	ProtocolFilter        []string           `json:"protocolFilter,omitempty"`
	NetworkIdFilter       string             `json:"networkIdFilter,omitempty"`
}

type GetACLRuleRequest struct {
	SiteID string `json:"-"`
	RuleID string `json:"-"`
}

type UpdateACLRuleRequest struct {
	SiteID                string             `json:"-"`
	RuleID                string             `json:"-"`
	Type                  string             `json:"type"`
	Enabled               bool               `json:"enabled"`
	Name                  string             `json:"name"`
	Description           string             `json:"description,omitempty"`
	Action                string             `json:"action"`
	EnforcingDeviceFilter *ACLDeviceFilter   `json:"enforcingDeviceFilter,omitempty"`
	SourceFilter          *ACLEndpointFilter `json:"sourceFilter,omitempty"`
	DestinationFilter     *ACLEndpointFilter `json:"destinationFilter,omitempty"`
	ProtocolFilter        []string           `json:"protocolFilter,omitempty"`
	NetworkIdFilter       string             `json:"networkIdFilter,omitempty"`
}

type DeleteACLRuleRequest struct {
	SiteID string `json:"-"`
	RuleID string `json:"-"`
}

type ACLRuleOrdering struct {
	RuleIDs []string `json:"ruleIds"`
}

type GetACLRuleOrderingRequest struct {
	SiteID string `json:"-"`
}

type UpdateACLRuleOrderingRequest struct {
	SiteID  string   `json:"-"`
	RuleIDs []string `json:"ruleIds"`
}

type DNSPolicy struct {
	Type             string          `json:"type"`
	ID               string          `json:"id,omitempty"`
	Enabled          bool            `json:"enabled"`
	Metadata         *EntityMetadata `json:"metadata,omitempty"`
	Domain           string          `json:"domain,omitempty"`
	IPv4Address      string          `json:"ipv4Address,omitempty"`
	IPv6Address      string          `json:"ipv6Address,omitempty"`
	TargetDomain     string          `json:"targetDomain,omitempty"`
	MailServerDomain string          `json:"mailServerDomain,omitempty"`
	Priority         *int            `json:"priority,omitempty"`
	Text             string          `json:"text,omitempty"`
	ServerDomain     string          `json:"serverDomain,omitempty"`
	Service          string          `json:"service,omitempty"`
	Protocol         string          `json:"protocol,omitempty"`
	Port             *int            `json:"port,omitempty"`
	Weight           *int            `json:"weight,omitempty"`
	IPAddress        string          `json:"ipAddress,omitempty"`
	TTLSeconds       *int            `json:"ttlSeconds,omitempty"`
}

type ListDNSPoliciesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateDNSPolicyRequest struct {
	SiteID           string `json:"-"`
	Type             string `json:"type"`
	Enabled          bool   `json:"enabled"`
	Domain           string `json:"domain,omitempty"`
	IPv4Address      string `json:"ipv4Address,omitempty"`
	IPv6Address      string `json:"ipv6Address,omitempty"`
	TargetDomain     string `json:"targetDomain,omitempty"`
	MailServerDomain string `json:"mailServerDomain,omitempty"`
	Priority         *int   `json:"priority,omitempty"`
	Text             string `json:"text,omitempty"`
	ServerDomain     string `json:"serverDomain,omitempty"`
	Service          string `json:"service,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	Port             *int   `json:"port,omitempty"`
	Weight           *int   `json:"weight,omitempty"`
	IPAddress        string `json:"ipAddress,omitempty"`
	TTLSeconds       *int   `json:"ttlSeconds,omitempty"`
}

type GetDNSPolicyRequest struct {
	SiteID   string `json:"-"`
	PolicyID string `json:"-"`
}

type UpdateDNSPolicyRequest struct {
	SiteID           string `json:"-"`
	PolicyID         string `json:"-"`
	Type             string `json:"type"`
	Enabled          bool   `json:"enabled"`
	Domain           string `json:"domain,omitempty"`
	IPv4Address      string `json:"ipv4Address,omitempty"`
	IPv6Address      string `json:"ipv6Address,omitempty"`
	TargetDomain     string `json:"targetDomain,omitempty"`
	MailServerDomain string `json:"mailServerDomain,omitempty"`
	Priority         *int   `json:"priority,omitempty"`
	Text             string `json:"text,omitempty"`
	ServerDomain     string `json:"serverDomain,omitempty"`
	Service          string `json:"service,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	Port             *int   `json:"port,omitempty"`
	Weight           *int   `json:"weight,omitempty"`
	IPAddress        string `json:"ipAddress,omitempty"`
	TTLSeconds       *int   `json:"ttlSeconds,omitempty"`
}

type DeleteDNSPolicyRequest struct {
	SiteID   string `json:"-"`
	PolicyID string `json:"-"`
}

type TrafficMatchingList struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name"`
	Type     string          `json:"type,omitempty"`
	Entries  []string        `json:"entries,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type ListTrafficMatchingListsRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type CreateTrafficMatchingListRequest struct {
	SiteID  string   `json:"-"`
	Name    string   `json:"name"`
	Type    string   `json:"type,omitempty"`
	Entries []string `json:"entries,omitempty"`
}

type GetTrafficMatchingListRequest struct {
	SiteID string `json:"-"`
	ListID string `json:"-"`
}

type UpdateTrafficMatchingListRequest struct {
	SiteID  string   `json:"-"`
	ListID  string   `json:"-"`
	Name    string   `json:"name"`
	Type    string   `json:"type,omitempty"`
	Entries []string `json:"entries,omitempty"`
}

type DeleteTrafficMatchingListRequest struct {
	SiteID string `json:"-"`
	ListID string `json:"-"`
}

type ListWANInterfacesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type ListVPNTunnelsRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type ListVPNServersRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type ListRadiusProfilesRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type ListDeviceTagsRequest struct {
	SiteID     string            `json:"-"`
	Pagination *PaginationParams `json:"-"`
}

type ListDPICategoriesRequest struct {
	SiteID string `json:"-"`
}

type ListDPIApplicationsRequest struct {
	SiteID string `json:"-"`
}

type ListCountriesRequest struct {
	SiteID string `json:"-"`
}

type WANInterface struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	Enabled  bool            `json:"enabled,omitempty"`
	Type     string          `json:"type,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type VPNTunnel struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	Enabled  bool            `json:"enabled,omitempty"`
	Type     string          `json:"type,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type VPNServer struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	Enabled  bool            `json:"enabled,omitempty"`
	Type     string          `json:"type,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type RadiusProfile struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type DeviceTag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DPIApplicationCategory struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DPIApplication struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryID string `json:"categoryId,omitempty"`
}

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}
