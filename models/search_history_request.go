package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var defaultStreamType = "logs"

// checks if the SearchHistoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHistoryRequest{}

// SearchHistoryRequest struct for SearchHistoryRequest
type SearchHistoryRequest struct {
	EndTime    int64          `json:"end_time"`
	Size       *int64         `json:"size,omitempty"`
	StartTime  int64          `json:"start_time"`
	StreamName NullableString `json:"stream_name,omitempty"`
	StreamType NullableString `json:"stream_type,omitempty"`
	TraceId    NullableString `json:"trace_id,omitempty"`
}

type _SearchHistoryRequest SearchHistoryRequest

// NewSearchHistoryRequest instantiates a new SearchHistoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHistoryRequest(startTime int64, endTime int64, streamName string, streamType *string) *SearchHistoryRequest {
	this := SearchHistoryRequest{}
	this.EndTime = EnsureMicro(endTime)
	this.StartTime = EnsureMicro(startTime)
	this.Size = &defaultSize
	this.StreamName.Set(&streamName)
	if streamType != nil {
		this.StreamType = *NewNullableString(streamType)
	} else {
		this.StreamType = *NewNullableString(&defaultStreamType)
	}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *SearchHistoryRequest) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SearchHistoryRequest) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SearchHistoryRequest) SetEndTime(v int64) {
	o.EndTime = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SearchHistoryRequest) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHistoryRequest) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SearchHistoryRequest) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SearchHistoryRequest) SetSize(v int64) {
	o.Size = &v
}

// GetStartTime returns the StartTime field value
func (o *SearchHistoryRequest) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SearchHistoryRequest) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SearchHistoryRequest) SetStartTime(v int64) {
	o.StartTime = v
}

// GetStreamName returns the StreamName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchHistoryRequest) GetStreamName() string {
	if o == nil || IsNil(o.StreamName.Get()) {
		var ret string
		return ret
	}
	return *o.StreamName.Get()
}

// GetStreamNameOk returns a tuple with the StreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchHistoryRequest) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamName.Get(), o.StreamName.IsSet()
}

// HasStreamName returns a boolean if a field has been set.
func (o *SearchHistoryRequest) HasStreamName() bool {
	if o != nil && o.StreamName.IsSet() {
		return true
	}

	return false
}

// SetStreamName gets a reference to the given NullableString and assigns it to the StreamName field.
func (o *SearchHistoryRequest) SetStreamName(v string) {
	o.StreamName.Set(&v)
}

// SetStreamNameNil sets the value for StreamName to be an explicit nil
func (o *SearchHistoryRequest) SetStreamNameNil() {
	o.StreamName.Set(nil)
}

// UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil
func (o *SearchHistoryRequest) UnsetStreamName() {
	o.StreamName.Unset()
}

// GetStreamType returns the StreamType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchHistoryRequest) GetStreamType() string {
	if o == nil || IsNil(o.StreamType.Get()) {
		var ret string
		return ret
	}
	return *o.StreamType.Get()
}

// GetStreamTypeOk returns a tuple with the StreamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchHistoryRequest) GetStreamTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamType.Get(), o.StreamType.IsSet()
}

// HasStreamType returns a boolean if a field has been set.
func (o *SearchHistoryRequest) HasStreamType() bool {
	if o != nil && o.StreamType.IsSet() {
		return true
	}

	return false
}

// SetStreamType gets a reference to the given NullableString and assigns it to the StreamType field.
func (o *SearchHistoryRequest) SetStreamType(v string) {
	o.StreamType.Set(&v)
}

// SetStreamTypeNil sets the value for StreamType to be an explicit nil
func (o *SearchHistoryRequest) SetStreamTypeNil() {
	o.StreamType.Set(nil)
}

// UnsetStreamType ensures that no value is present for StreamType, not even an explicit nil
func (o *SearchHistoryRequest) UnsetStreamType() {
	o.StreamType.Unset()
}

// GetTraceId returns the TraceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchHistoryRequest) GetTraceId() string {
	if o == nil || IsNil(o.TraceId.Get()) {
		var ret string
		return ret
	}
	return *o.TraceId.Get()
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchHistoryRequest) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceId.Get(), o.TraceId.IsSet()
}

// HasTraceId returns a boolean if a field has been set.
func (o *SearchHistoryRequest) HasTraceId() bool {
	if o != nil && o.TraceId.IsSet() {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given NullableString and assigns it to the TraceId field.
func (o *SearchHistoryRequest) SetTraceId(v string) {
	o.TraceId.Set(&v)
}

// SetTraceIdNil sets the value for TraceId to be an explicit nil
func (o *SearchHistoryRequest) SetTraceIdNil() {
	o.TraceId.Set(nil)
}

// UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil
func (o *SearchHistoryRequest) UnsetTraceId() {
	o.TraceId.Unset()
}

func (o SearchHistoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHistoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end_time"] = o.EndTime
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["start_time"] = o.StartTime
	if o.StreamName.IsSet() {
		toSerialize["stream_name"] = o.StreamName.Get()
	}
	if o.StreamType.IsSet() {
		toSerialize["stream_type"] = o.StreamType.Get()
	}
	if o.TraceId.IsSet() {
		toSerialize["trace_id"] = o.TraceId.Get()
	}
	return toSerialize, nil
}

func (o *SearchHistoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end_time",
		"start_time",
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

	varSearchHistoryRequest := _SearchHistoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchHistoryRequest)

	if err != nil {
		return err
	}

	*o = SearchHistoryRequest(varSearchHistoryRequest)

	return err
}

type NullableSearchHistoryRequest struct {
	value *SearchHistoryRequest
	isSet bool
}

func (v NullableSearchHistoryRequest) Get() *SearchHistoryRequest {
	return v.value
}

func (v *NullableSearchHistoryRequest) Set(val *SearchHistoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHistoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHistoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHistoryRequest(val *SearchHistoryRequest) *NullableSearchHistoryRequest {
	return &NullableSearchHistoryRequest{value: val, isSet: true}
}

func (v NullableSearchHistoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHistoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
