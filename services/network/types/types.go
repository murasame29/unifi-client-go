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

type ApplicationInfo struct {
	ApplicationVersion string `json:"applicationVersion"`
}

type Site struct {
	ID                string `json:"id"`
	InternalReference string `json:"internalReference"`
	Name              string `json:"name"`
}

type AdoptedDevice struct {
	ID              string          `json:"id"`
	MacAddress      string          `json:"macAddress"`
	IPAddress       string          `json:"ipAddress"`
	Name            string          `json:"name"`
	Model           string          `json:"model"`
	Supported       bool            `json:"supported"`
	State           string          `json:"state"`
	FirmwareVersion string          `json:"firmwareVersion,omitempty"`
	FirmwareUpdatable bool          `json:"firmwareUpdatable"`
	AdoptedAt       string          `json:"adoptedAt,omitempty"`
	ProvisionedAt   string          `json:"provisionedAt,omitempty"`
	ConfigurationID string          `json:"configurationId"`
	Uplink          json.RawMessage `json:"uplink,omitempty"`
	Features        json.RawMessage `json:"features,omitempty"`
	Interfaces      json.RawMessage `json:"interfaces,omitempty"`
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

type DeviceStatistics struct {
	json.RawMessage
}

type PendingDevice struct {
	MacAddress string `json:"macAddress"`
	IPAddress  string `json:"ipAddress"`
	Model      string `json:"model"`
	Supported  bool   `json:"supported"`
}

type ConnectedClient struct {
	ID         string          `json:"id"`
	MacAddress string          `json:"macAddress"`
	IPAddress  string          `json:"ipAddress"`
	Name       string          `json:"name,omitempty"`
	Type       string          `json:"type"`
	Hostname   string          `json:"hostname,omitempty"`
	Connected  bool            `json:"connected"`
	Extra      json.RawMessage `json:"extra,omitempty"`
}

type ClientActionRequest struct {
	Action string `json:"action"`
}

type Network struct {
	Management   string          `json:"management"`
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Enabled      bool            `json:"enabled"`
	VlanID       int             `json:"vlanId"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
	DhcpGuarding json.RawMessage `json:"dhcpGuarding,omitempty"`
	Default      bool            `json:"default,omitempty"`
}

type CreateNetworkRequest struct {
	Management   string          `json:"management"`
	Name         string          `json:"name"`
	Enabled      bool            `json:"enabled"`
	VlanID       int             `json:"vlanId"`
	DhcpGuarding json.RawMessage `json:"dhcpGuarding,omitempty"`
}

type UpdateNetworkRequest struct {
	Management   string          `json:"management"`
	Name         string          `json:"name"`
	Enabled      bool            `json:"enabled"`
	VlanID       int             `json:"vlanId"`
	DhcpGuarding json.RawMessage `json:"dhcpGuarding,omitempty"`
}

type NetworkReference struct {
	json.RawMessage
}

type WifiBroadcast struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name"`
	Enabled  bool            `json:"enabled"`
	Security json.RawMessage `json:"security,omitempty"`
	Extra    json.RawMessage `json:"extra,omitempty"`
}

type CreateWifiBroadcastRequest struct {
	Name     string          `json:"name"`
	Enabled  bool            `json:"enabled"`
	Security json.RawMessage `json:"security,omitempty"`
}

type UpdateWifiBroadcastRequest struct {
	Name     string          `json:"name"`
	Enabled  bool            `json:"enabled"`
	Security json.RawMessage `json:"security,omitempty"`
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
	Count    int    `json:"count"`
	Duration int    `json:"duration"`
	DataQuota int   `json:"dataQuota,omitempty"`
	Note     string `json:"note,omitempty"`
}

type DeleteVouchersRequest struct {
	IDs []string `json:"ids"`
}

type FirewallZone struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type CreateFirewallZoneRequest struct {
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type UpdateFirewallZoneRequest struct {
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type FirewallPolicy struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type CreateFirewallPolicyRequest struct {
	json.RawMessage
}

type UpdateFirewallPolicyRequest struct {
	json.RawMessage
}

type PatchFirewallPolicyRequest struct {
	json.RawMessage
}

type FirewallPolicyOrdering struct {
	json.RawMessage
}

type UpdateFirewallPolicyOrderingRequest struct {
	json.RawMessage
}

type ACLRule struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type CreateACLRuleRequest struct {
	json.RawMessage
}

type UpdateACLRuleRequest struct {
	json.RawMessage
}

type ACLRuleOrdering struct {
	json.RawMessage
}

type UpdateACLRuleOrderingRequest struct {
	json.RawMessage
}

type DNSPolicy struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type CreateDNSPolicyRequest struct {
	json.RawMessage
}

type UpdateDNSPolicyRequest struct {
	json.RawMessage
}

type TrafficMatchingList struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name"`
	Extra json.RawMessage `json:"extra,omitempty"`
}

type CreateTrafficMatchingListRequest struct {
	json.RawMessage
}

type UpdateTrafficMatchingListRequest struct {
	json.RawMessage
}

type WANInterface struct {
	json.RawMessage
}

type VPNTunnel struct {
	json.RawMessage
}

type VPNServer struct {
	json.RawMessage
}

type RadiusProfile struct {
	json.RawMessage
}

type DeviceTag struct {
	json.RawMessage
}

type DPIApplicationCategory struct {
	json.RawMessage
}

type DPIApplication struct {
	json.RawMessage
}

type Country struct {
	json.RawMessage
}
