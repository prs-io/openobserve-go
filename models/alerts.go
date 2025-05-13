package models

type FrequencyType string

const (
	FrequencyTypeCron    FrequencyType = "cron"
	FrequencyTypeMinutes FrequencyType = "minutes"
)

type TriggerCondition struct {
	Cron            *string       `json:"cron,omitempty"`
	Frequency       *int64        `json:"frequency,omitempty"`
	FrequencyType   FrequencyType `json:"frequency_type"`
	Operator        Operator      `json:"operator"`
	Period          int64         `json:"period"`
	Silence         *int64        `json:"silence,omitempty"`
	Threshold       *int64        `json:"threshold,omitempty"`
	Timezone        *string       `json:"timezone,omitempty"`
	ToleranceInSecs *int64        `json:"tolerance_in_secs,omitempty"`
}

type Alert struct {
	ContextAttributes map[string]string `json:"context_attributes,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Destinations      []string          `json:"destinations"`
	Enabled           *bool             `json:"enabled,omitempty"`
	ID                string            `json:"id"`
	IsRealTime        *bool             `json:"is_real_time,omitempty"`
	LastEditedBy      *string           `json:"last_edited_by,omitempty"`
	LastSatisfiedAt   *int64            `json:"last_satisfied_at,omitempty"`
	LastTriggeredAt   *int64            `json:"last_triggered_at,omitempty"`
	Name              string            `json:"name"`
	OrgID             string            `json:"org_id"`
	Owner             *string           `json:"owner,omitempty"`
	QueryCondition    QueryCondition    `json:"query_condition"`
	RowTemplate       *string           `json:"row_template,omitempty"`
	StreamName        *string           `json:"stream_name,omitempty"`
	StreamType        *StreamType       `json:"stream_type,omitempty"`
	TriggerCondition  *TriggerCondition `json:"trigger_condition,omitempty"`
	TzOffset          *int32            `json:"tz_offset,omitempty"`
	UpdatedAt         *int64            `json:"updated_at,omitempty"`
}

type ListAlertsResponse struct {
	Alerts []ListAlertsResponseItem `json:"list"`
}

type ListAlertsResponseItem struct {
	AlertID         string         `json:"alert_id"`
	Condition       QueryCondition `json:"condition"`
	Description     *string        `json:"description,omitempty"`
	FolderID        string         `json:"folder_id"`
	FolderName      string         `json:"folder_name"`
	Name            string         `json:"name"`
	Owner           *string        `json:"owner,omitempty"`
	LastSatisfiedAt *int64         `json:"last_satisfied_at,omitempty"`
	LastTriggeredAt *int64         `json:"last_triggered_at,omitempty"`
}

type QueryType string

const (
	QueryTypeCustom QueryType = "custom"
	QueryTypeSQL    QueryType = "sql"
	QueryTypePromQL QueryType = "promql"
)

type QueryCondition struct {
	Aggregation     *Aggregation           `json:"aggregation,omitempty"`
	conditions      []Condition            `json:"conditions,omitempty"`
	MultiTimeRange  []*CompareHistoricData `json:"multi_time_range,omitempty"`
	PromQL          *string                `json:"promql,omitempty"`
	PromQLCondition *PromQLCondition       `json:"promql_condition,omitempty"`
	SearchEventType *SearchEventType       `json:"search_event_type,omitempty"`
	SQL             *string                `json:"sql,omitempty"`
	Type            QueryType              `json:"type"`
	VRLFunction     *string                `json:"vrl_function,omitempty"`

	// GroupBy []string   `json:"group_by,omitempty"`
	// Having  *Condition `json:"having,omitempty"`
}

type Aggregation struct {
	Function string `json:"function"` // Enum: avg, min, max, sum, count, median, p50, p75, p90, p95, p99
}

type Operator string

const (
	OperatorEQ          Operator = "="
	OperatorNEQ         Operator = "!="
	OperatorGT          Operator = ">"
	OperatorGTEQ        Operator = ">="
	OperatorLT          Operator = "<"
	OperatorLTEQ        Operator = "<="
	OperatorContains    Operator = "Contains"
	OperatorNotContains Operator = "NotContains"
)

type Condition struct {
	Column     string      `json:"column"`
	IgnoreCase bool        `json:"ignore_case"`
	Operator   Operator    `json:"operator"` // Enum: =, !=, >, >=, <, <=, Contains, NotContains
	Value      interface{} `json:"value"`
}

type CompareHistoricData struct {
	OffSet string `json:"offSet"`
}

type PromQLCondition struct {
	Column     string      `json:"column"`
	IgnoreCase bool        `json:"ignore_case"`
	Operator   Operator    `json:"operator"`
	Value      interface{} `json:"value"`
}
