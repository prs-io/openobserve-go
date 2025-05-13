package models

// DashboardsResponse represents the response body for listing dashboards.
type DashboardsResponse struct {
	Dashboards []DashboardsResponseItem `json:"dashboards"`
}

// ListDashboardsResponseBodyItem represents an item in the dashboards list.
type DashboardsResponseItem struct {
	Created     string      `json:"created"` // date-time string
	DashboardID string      `json:"dashboard_id"`
	Description string      `json:"description"`
	FolderID    string      `json:"folder_id"`
	FolderName  string      `json:"folder_name"`
	Hash        string      `json:"hash"`
	Owner       string      `json:"owner"`
	Role        string      `json:"role"`
	Title       string      `json:"title"`
	V1          interface{} `json:"v1,omitempty"`
	V2          interface{} `json:"v2,omitempty"`
	V3          interface{} `json:"v3,omitempty"`
	V4          interface{} `json:"v4,omitempty"`
	V5          interface{} `json:"v5,omitempty"`
	Version     int32       `json:"version"`
}

type DashboardResponse struct {
	V1      interface{} `json:"v1,omitempty"`
	V2      interface{} `json:"v2,omitempty"`
	V3      interface{} `json:"v3,omitempty"`
	V4      interface{} `json:"v4,omitempty"`
	V5      interface{} `json:"v5,omitempty"`
	Version int32       `json:"version"`
	Hash    string      `json:"hash"`
}
