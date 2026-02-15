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

type AdoptDeviceRequest struct {
	MacAddress        string `json:"macAddress"`
	IgnoreDeviceLimit bool   `json:"ignoreDeviceLimit"`
}

type PortActionRequest struct {
	Action string `json:"action"`
}

type DeviceActionRequest struct {
	Action string `json:"action"`
}

type DeviceTemperature struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type DeviceStatistics struct {
	Temperatures []DeviceTemperature `json:"temperatures,omitempty"`
	Uptime       int                 `json:"uptime,omitempty"`
	LoadAverage1 float64             `json:"loadAverage1,omitempty"`
	LoadAverage5 float64             `json:"loadAverage5,omitempty"`
}

type PendingDevice struct {
	MacAddress string `json:"macAddress"`
	IPAddress  string `json:"ipAddress"`
	Model      string `json:"model"`
	Supported  bool   `json:"supported"`
}

type ConnectedClient struct {
	ID         string `json:"id"`
	MacAddress string `json:"macAddress"`
	IPAddress  string `json:"ipAddress"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type"`
	Hostname   string `json:"hostname,omitempty"`
	Connected  bool   `json:"connected"`
}

type ClientActionRequest struct {
	Action string `json:"action"`
}

type DhcpGuarding struct {
	TrustedDhcpServerIPAddresses []string `json:"trustedDhcpServerIpAddresses,omitempty"`
}

type Network struct {
	Management   string          `json:"management"`
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Enabled      bool            `json:"enabled"`
	VlanID       int             `json:"vlanId"`
	Metadata     *EntityMetadata `json:"metadata,omitempty"`
	DhcpGuarding *DhcpGuarding   `json:"dhcpGuarding,omitempty"`
	Default      bool            `json:"default,omitempty"`
}

type CreateNetworkRequest struct {
	Management   string        `json:"management"`
	Name         string        `json:"name"`
	Enabled      bool          `json:"enabled"`
	VlanID       int           `json:"vlanId"`
	DhcpGuarding *DhcpGuarding `json:"dhcpGuarding,omitempty"`
}

type UpdateNetworkRequest struct {
	Management   string        `json:"management"`
	Name         string        `json:"name"`
	Enabled      bool          `json:"enabled"`
	VlanID       int           `json:"vlanId"`
	DhcpGuarding *DhcpGuarding `json:"dhcpGuarding,omitempty"`
}

type NetworkReference struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type WifiNetworkReference struct {
	Type      string `json:"type,omitempty"`
	NetworkID string `json:"networkId,omitempty"`
}

type WifiSecurityConfiguration struct {
	Type                string                   `json:"type"`
	RadiusConfiguration *WifiRadiusConfiguration `json:"radiusConfiguration,omitempty"`
	Password            string                   `json:"password,omitempty"`
	SaeTransitionMode   bool                     `json:"saeTransitionMode,omitempty"`
	Protocols           *WifiSecurityProtocols   `json:"protocols,omitempty"`
}

type WifiRadiusConfiguration struct {
	RadiusProfileID string `json:"radiusProfileId,omitempty"`
}

type WifiSecurityProtocols struct {
	Rsn bool `json:"rsn,omitempty"`
	Wpa bool `json:"wpa,omitempty"`
}

type BroadcastingDeviceFilter struct {
	Type      string   `json:"type"`
	DeviceIDs []string `json:"deviceIds,omitempty"`
}

type MdnsProxyConfiguration struct {
	Mode string `json:"mode"`
}

type MulticastFilteringPolicy struct {
	Action string `json:"action"`
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
	Type string `json:"type"`
	Day  string `json:"day"`
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

type CreateWifiBroadcastRequest struct {
	Type                                string                            `json:"type"`
	Name                                string                            `json:"name"`
	Network                             *WifiNetworkReference             `json:"network,omitempty"`
	Enabled                             bool                              `json:"enabled"`
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

type UpdateWifiBroadcastRequest struct {
	Type                                string                            `json:"type"`
	Name                                string                            `json:"name"`
	Network                             *WifiNetworkReference             `json:"network,omitempty"`
	Enabled                             bool                              `json:"enabled"`
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

type Voucher struct {
	ID        string `json:"id,omitempty"`
	Code      string `json:"code,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	DataQuota int    `json:"dataQuota,omitempty"`
	Note      string `json:"note,omitempty"`
	Used      bool   `json:"used,omitempty"`
}

type GenerateVouchersRequest struct {
	Count     int    `json:"count"`
	Duration  int    `json:"duration"`
	DataQuota int    `json:"dataQuota,omitempty"`
	Note      string `json:"note,omitempty"`
}

type DeleteVouchersRequest struct {
	IDs []string `json:"ids"`
}

type FirewallZone struct {
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name"`
	NetworkIDs []string        `json:"networkIds,omitempty"`
	Metadata   *EntityMetadata `json:"metadata,omitempty"`
}

type CreateFirewallZoneRequest struct {
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds,omitempty"`
}

type UpdateFirewallZoneRequest struct {
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds,omitempty"`
}

type FirewallPolicyAction struct {
	Type               string `json:"type"`
	AllowReturnTraffic *bool  `json:"allowReturnTraffic,omitempty"`
}

type TrafficFilter struct {
	Type string `json:"type"`
}

type FirewallPolicyEndpoint struct {
	ZoneID        string         `json:"zoneId"`
	TrafficFilter *TrafficFilter `json:"trafficFilter,omitempty"`
}

type FirewallIPProtocolScope struct {
	IPVersion string `json:"ipVersion"`
}

type FirewallSchedule struct {
	Mode string `json:"mode"`
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

type CreateFirewallPolicyRequest struct {
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

type UpdateFirewallPolicyRequest struct {
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

type PatchFirewallPolicyRequest struct {
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

type UpdateFirewallPolicyOrderingRequest struct {
	PolicyIDs []string `json:"policyIds"`
}

type ACLDeviceFilter struct {
	Type      string   `json:"type"`
	DeviceIDs []string `json:"deviceIds,omitempty"`
}

type ACLEndpointFilter struct {
	NetworkID  string   `json:"networkId,omitempty"`
	IPAddress  string   `json:"ipAddress,omitempty"`
	PortRanges []string `json:"portRanges,omitempty"`
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
}

type CreateACLRuleRequest struct {
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
}

type UpdateACLRuleRequest struct {
	Type                  string             `json:"type"`
	Enabled               bool               `json:"enabled"`
	Name                  string             `json:"name"`
	Description           string             `json:"description,omitempty"`
	Action                string             `json:"action"`
	EnforcingDeviceFilter *ACLDeviceFilter   `json:"enforcingDeviceFilter,omitempty"`
	SourceFilter          *ACLEndpointFilter `json:"sourceFilter,omitempty"`
	DestinationFilter     *ACLEndpointFilter `json:"destinationFilter,omitempty"`
	ProtocolFilter        []string           `json:"protocolFilter,omitempty"`
}

type ACLRuleOrdering struct {
	RuleIDs []string `json:"ruleIds"`
}

type UpdateACLRuleOrderingRequest struct {
	RuleIDs []string `json:"ruleIds"`
}

type DNSPolicy struct {
	Type        string          `json:"type"`
	ID          string          `json:"id,omitempty"`
	Enabled     bool            `json:"enabled"`
	Metadata    *EntityMetadata `json:"metadata,omitempty"`
	Domain      string          `json:"domain,omitempty"`
	IPv4Address string          `json:"ipv4Address,omitempty"`
	TTLSeconds  int             `json:"ttlSeconds,omitempty"`
}

type CreateDNSPolicyRequest struct {
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	Domain      string `json:"domain,omitempty"`
	IPv4Address string `json:"ipv4Address,omitempty"`
	TTLSeconds  int    `json:"ttlSeconds,omitempty"`
}

type UpdateDNSPolicyRequest struct {
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	Domain      string `json:"domain,omitempty"`
	IPv4Address string `json:"ipv4Address,omitempty"`
	TTLSeconds  int    `json:"ttlSeconds,omitempty"`
}

type TrafficMatchingList struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name"`
	Type     string          `json:"type,omitempty"`
	Entries  []string        `json:"entries,omitempty"`
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

type CreateTrafficMatchingListRequest struct {
	Name    string   `json:"name"`
	Type    string   `json:"type,omitempty"`
	Entries []string `json:"entries,omitempty"`
}

type UpdateTrafficMatchingListRequest struct {
	Name    string   `json:"name"`
	Type    string   `json:"type,omitempty"`
	Entries []string `json:"entries,omitempty"`
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
