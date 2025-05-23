package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var defaultSize = int64(50)

// checks if the SearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchQuery{}

// SearchQuery struct for SearchQuery
type SearchQuery struct {
	ActionId        NullableString `json:"action_id,omitempty"`
	EndTime         int64          `json:"end_time"`
	From            *int64         `json:"from,omitempty"`
	QueryFn         NullableString `json:"query_fn,omitempty"`
	QueryType       *string        `json:"query_type,omitempty"`
	QuickMode       *bool          `json:"quick_mode,omitempty"`
	Size            *int64         `json:"size,omitempty"`
	SkipWal         *bool          `json:"skip_wal,omitempty"`
	Sql             string         `json:"sql"`
	StartTime       int64          `json:"start_time"`
	StreamingId     NullableString `json:"streaming_id,omitempty"`
	StreamingOutput *bool          `json:"streaming_output,omitempty"`
	TrackTotalHits  *bool          `json:"track_total_hits,omitempty"`
	UsesZoFn        *bool          `json:"uses_zo_fn,omitempty"`
}

type _SearchQuery SearchQuery

// NewSearchQuery instantiates a new SearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchQuery(startTime, endTime int64, sql string, size *int64) *SearchQuery {
	this := SearchQuery{}
	//True := true
	from := int64(0)
	this.Size = &defaultSize
	//this.QuickMode = &True
	this.From = &from
	this.Sql = sql
	this.StartTime = EnsureMicro(startTime)
	this.EndTime = EnsureMicro(endTime)
	if size != nil {
		if *size > 0 {
			this.Size = size
		}
	}
	return &this
}

// GetActionId returns the ActionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetActionId() string {
	if o == nil || IsNil(o.ActionId.Get()) {
		var ret string
		return ret
	}
	return *o.ActionId.Get()
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionId.Get(), o.ActionId.IsSet()
}

// HasActionId returns a boolean if a field has been set.
func (o *SearchQuery) HasActionId() bool {
	if o != nil && o.ActionId.IsSet() {
		return true
	}

	return false
}

// SetActionId gets a reference to the given NullableString and assigns it to the ActionId field.
func (o *SearchQuery) SetActionId(v string) {
	o.ActionId.Set(&v)
}

// SetActionIdNil sets the value for ActionId to be an explicit nil
func (o *SearchQuery) SetActionIdNil() {
	o.ActionId.Set(nil)
}

// UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
func (o *SearchQuery) UnsetActionId() {
	o.ActionId.Unset()
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SearchQuery) GetFrom() int64 {
	if o == nil || IsNil(o.From) {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetFromOk() (*int64, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SearchQuery) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *SearchQuery) SetFrom(v int64) {
	o.From = &v
}

// GetQueryFn returns the QueryFn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetQueryFn() string {
	if o == nil || IsNil(o.QueryFn.Get()) {
		var ret string
		return ret
	}
	return *o.QueryFn.Get()
}

// GetQueryFnOk returns a tuple with the QueryFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetQueryFnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryFn.Get(), o.QueryFn.IsSet()
}

// HasQueryFn returns a boolean if a field has been set.
func (o *SearchQuery) HasQueryFn() bool {
	if o != nil && o.QueryFn.IsSet() {
		return true
	}

	return false
}

// SetQueryFn gets a reference to the given NullableString and assigns it to the QueryFn field.
func (o *SearchQuery) SetQueryFn(v string) {
	o.QueryFn.Set(&v)
}

// SetQueryFnNil sets the value for QueryFn to be an explicit nil
func (o *SearchQuery) SetQueryFnNil() {
	o.QueryFn.Set(nil)
}

// UnsetQueryFn ensures that no value is present for QueryFn, not even an explicit nil
func (o *SearchQuery) UnsetQueryFn() {
	o.QueryFn.Unset()
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *SearchQuery) GetQueryType() string {
	if o == nil || IsNil(o.QueryType) {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetQueryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryType) {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *SearchQuery) HasQueryType() bool {
	if o != nil && !IsNil(o.QueryType) {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *SearchQuery) SetQueryType(v string) {
	o.QueryType = &v
}

// GetQuickMode returns the QuickMode field value if set, zero value otherwise.
func (o *SearchQuery) GetQuickMode() bool {
	if o == nil || IsNil(o.QuickMode) {
		var ret bool
		return ret
	}
	return *o.QuickMode
}

// GetQuickModeOk returns a tuple with the QuickMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetQuickModeOk() (*bool, bool) {
	if o == nil || IsNil(o.QuickMode) {
		return nil, false
	}
	return o.QuickMode, true
}

// HasQuickMode returns a boolean if a field has been set.
func (o *SearchQuery) HasQuickMode() bool {
	if o != nil && !IsNil(o.QuickMode) {
		return true
	}

	return false
}

// SetQuickMode gets a reference to the given bool and assigns it to the QuickMode field.
func (o *SearchQuery) SetQuickMode(v bool) {
	o.QuickMode = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SearchQuery) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SearchQuery) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SearchQuery) SetSize(v int64) {
	o.Size = &v
}

// GetSkipWal returns the SkipWal field value if set, zero value otherwise.
func (o *SearchQuery) GetSkipWal() bool {
	if o == nil || IsNil(o.SkipWal) {
		var ret bool
		return ret
	}
	return *o.SkipWal
}

// GetSkipWalOk returns a tuple with the SkipWal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetSkipWalOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWal) {
		return nil, false
	}
	return o.SkipWal, true
}

// HasSkipWal returns a boolean if a field has been set.
func (o *SearchQuery) HasSkipWal() bool {
	if o != nil && !IsNil(o.SkipWal) {
		return true
	}

	return false
}

// SetSkipWal gets a reference to the given bool and assigns it to the SkipWal field.
func (o *SearchQuery) SetSkipWal(v bool) {
	o.SkipWal = &v
}

// GetSql returns the Sql field value
func (o *SearchQuery) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *SearchQuery) SetSql(v string) {
	o.Sql = v
}

// GetStreamingId returns the StreamingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetStreamingId() string {
	if o == nil || IsNil(o.StreamingId.Get()) {
		var ret string
		return ret
	}
	return *o.StreamingId.Get()
}

// GetStreamingIdOk returns a tuple with the StreamingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetStreamingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamingId.Get(), o.StreamingId.IsSet()
}

// HasStreamingId returns a boolean if a field has been set.
func (o *SearchQuery) HasStreamingId() bool {
	if o != nil && o.StreamingId.IsSet() {
		return true
	}

	return false
}

// SetStreamingId gets a reference to the given NullableString and assigns it to the StreamingId field.
func (o *SearchQuery) SetStreamingId(v string) {
	o.StreamingId.Set(&v)
}

// SetStreamingIdNil sets the value for StreamingId to be an explicit nil
func (o *SearchQuery) SetStreamingIdNil() {
	o.StreamingId.Set(nil)
}

// UnsetStreamingId ensures that no value is present for StreamingId, not even an explicit nil
func (o *SearchQuery) UnsetStreamingId() {
	o.StreamingId.Unset()
}

// GetStreamingOutput returns the StreamingOutput field value if set, zero value otherwise.
func (o *SearchQuery) GetStreamingOutput() bool {
	if o == nil || IsNil(o.StreamingOutput) {
		var ret bool
		return ret
	}
	return *o.StreamingOutput
}

// GetStreamingOutputOk returns a tuple with the StreamingOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetStreamingOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.StreamingOutput) {
		return nil, false
	}
	return o.StreamingOutput, true
}

// HasStreamingOutput returns a boolean if a field has been set.
func (o *SearchQuery) HasStreamingOutput() bool {
	if o != nil && !IsNil(o.StreamingOutput) {
		return true
	}

	return false
}

// SetStreamingOutput gets a reference to the given bool and assigns it to the StreamingOutput field.
func (o *SearchQuery) SetStreamingOutput(v bool) {
	o.StreamingOutput = &v
}

// GetTrackTotalHits returns the TrackTotalHits field value if set, zero value otherwise.
func (o *SearchQuery) GetTrackTotalHits() bool {
	if o == nil || IsNil(o.TrackTotalHits) {
		var ret bool
		return ret
	}
	return *o.TrackTotalHits
}

// GetTrackTotalHitsOk returns a tuple with the TrackTotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetTrackTotalHitsOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackTotalHits) {
		return nil, false
	}
	return o.TrackTotalHits, true
}

// HasTrackTotalHits returns a boolean if a field has been set.
func (o *SearchQuery) HasTrackTotalHits() bool {
	if o != nil && !IsNil(o.TrackTotalHits) {
		return true
	}

	return false
}

// SetTrackTotalHits gets a reference to the given bool and assigns it to the TrackTotalHits field.
func (o *SearchQuery) SetTrackTotalHits(v bool) {
	o.TrackTotalHits = &v
}

// GetUsesZoFn returns the UsesZoFn field value if set, zero value otherwise.
func (o *SearchQuery) GetUsesZoFn() bool {
	if o == nil || IsNil(o.UsesZoFn) {
		var ret bool
		return ret
	}
	return *o.UsesZoFn
}

// GetUsesZoFnOk returns a tuple with the UsesZoFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetUsesZoFnOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesZoFn) {
		return nil, false
	}
	return o.UsesZoFn, true
}

// HasUsesZoFn returns a boolean if a field has been set.
func (o *SearchQuery) HasUsesZoFn() bool {
	if o != nil && !IsNil(o.UsesZoFn) {
		return true
	}

	return false
}

// SetUsesZoFn gets a reference to the given bool and assigns it to the UsesZoFn field.
func (o *SearchQuery) SetUsesZoFn(v bool) {
	o.UsesZoFn = &v
}

func (o SearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionId.IsSet() {
		toSerialize["action_id"] = o.ActionId.Get()
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if o.QueryFn.IsSet() {
		toSerialize["query_fn"] = o.QueryFn.Get()
	}
	if !IsNil(o.QueryType) {
		toSerialize["query_type"] = o.QueryType
	}
	if !IsNil(o.QuickMode) {
		toSerialize["quick_mode"] = o.QuickMode
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SkipWal) {
		toSerialize["skip_wal"] = o.SkipWal
	}
	toSerialize["sql"] = o.Sql
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if o.StreamingId.IsSet() {
		toSerialize["streaming_id"] = o.StreamingId.Get()
	}
	if !IsNil(o.StreamingOutput) {
		toSerialize["streaming_output"] = o.StreamingOutput
	}
	if !IsNil(o.TrackTotalHits) {
		toSerialize["track_total_hits"] = o.TrackTotalHits
	}
	if !IsNil(o.UsesZoFn) {
		toSerialize["uses_zo_fn"] = o.UsesZoFn
	}
	return toSerialize, nil
}

func (o *SearchQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sql",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchQuery := _SearchQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchQuery)

	if err != nil {
		return err
	}

	*o = SearchQuery(varSearchQuery)

	return err
}

type NullableSearchQuery struct {
	value *SearchQuery
	isSet bool
}

func (v NullableSearchQuery) Get() *SearchQuery {
	return v.value
}

func (v *NullableSearchQuery) Set(val *SearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchQuery(val *SearchQuery) *NullableSearchQuery {
	return &NullableSearchQuery{value: val, isSet: true}
}

func (v NullableSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
