package types

import "encoding/json"

type PaginatedResponse[T any] struct {
	Code           json.RawMessage `json:"code,omitempty"`
	Data           []T             `json:"data"`
	HTTPStatusCode int             `json:"httpStatusCode"`
	TraceID        string          `json:"traceId"`
	NextToken      string          `json:"nextToken,omitempty"`
}

type SingleResponse[T any] struct {
	Data           T      `json:"data"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	TraceID        string `json:"traceId"`
}

type Host struct {
	ID                        string          `json:"id"`
	HardwareID                string          `json:"hardwareId"`
	Type                      string          `json:"type"`
	IPAddress                 string          `json:"ipAddress"`
	Owner                     bool            `json:"owner"`
	IsBlocked                 bool            `json:"isBlocked"`
	RegistrationTime          string          `json:"registrationTime"`
	LastConnectionStateChange string          `json:"lastConnectionStateChange"`
	LatestBackupTime          string          `json:"latestBackupTime"`
	UserData                  json.RawMessage `json:"userData,omitempty"`
	ReportedState             json.RawMessage `json:"reportedState,omitempty"`
}

type Site struct {
	SiteID     string          `json:"siteId"`
	HostID     string          `json:"hostId"`
	Meta       json.RawMessage `json:"meta,omitempty"`
	Statistics json.RawMessage `json:"statistics,omitempty"`
	Permission string          `json:"permission"`
	IsOwner    bool            `json:"isOwner"`
}

type DeviceHost struct {
	HostID    string   `json:"hostId"`
	HostName  string   `json:"hostName"`
	Devices   []Device `json:"devices"`
	UpdatedAt string   `json:"updatedAt"`
}

type Device struct {
	ID              string          `json:"id"`
	MAC             string          `json:"mac"`
	Name            string          `json:"name"`
	Model           string          `json:"model"`
	Shortname       string          `json:"shortname"`
	IP              string          `json:"ip"`
	ProductLine     string          `json:"productLine"`
	Status          string          `json:"status"`
	Version         string          `json:"version"`
	FirmwareStatus  string          `json:"firmwareStatus"`
	UpdateAvailable string          `json:"updateAvailable"`
	IsConsole       bool            `json:"isConsole"`
	IsManaged       bool            `json:"isManaged"`
	StartupTime     string          `json:"startupTime"`
	AdoptionTime    *string         `json:"adoptionTime"`
	Note            string          `json:"note"`
	UIDB            json.RawMessage `json:"uidb,omitempty"`
}

type ISPMetrics struct {
	MetricType string            `json:"metricType"`
	Periods    []ISPMetricPeriod `json:"periods"`
	HostID     string            `json:"hostId"`
	SiteID     string            `json:"siteId"`
}

type ISPMetricPeriod struct {
	Data       json.RawMessage `json:"data,omitempty"`
	MetricTime string          `json:"metricTime"`
	Version    string          `json:"version"`
}

type ISPMetricsQueryRequest struct {
	Sites []ISPMetricsQuerySite `json:"sites"`
}

type ISPMetricsQuerySite struct {
	BeginTimestamp string `json:"beginTimestamp"`
	HostID         string `json:"hostId"`
	EndTimestamp   string `json:"endTimestamp"`
	SiteID         string `json:"siteId"`
}

type ISPMetricsQueryResponse struct {
	Data           json.RawMessage `json:"data"`
	HTTPStatusCode int             `json:"httpStatusCode"`
	TraceID        string          `json:"traceId"`
}

type SDWANConfig struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
