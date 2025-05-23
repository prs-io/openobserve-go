package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SearchPartitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPartitionRequest{}

// SearchPartitionRequest struct for SearchPartitionRequest
type SearchPartitionRequest struct {
	Clusters        []string         `json:"clusters,omitempty"`
	Encoding        *RequestEncoding `json:"encoding,omitempty"`
	EndTime         int64            `json:"end_time"`
	QueryFn         NullableString   `json:"query_fn,omitempty"`
	Regions         []string         `json:"regions,omitempty"`
	Sql             string           `json:"sql"`
	StartTime       int64            `json:"start_time"`
	StreamingOutput *bool            `json:"streaming_output,omitempty"`
}

type _SearchPartitionRequest SearchPartitionRequest

// NewSearchPartitionRequest instantiates a new SearchPartitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPartitionRequest(startTime int64, endTime int64, sql string) *SearchPartitionRequest {
	this := SearchPartitionRequest{}
	this.EndTime = endTime
	this.Sql = sql
	this.StartTime = startTime
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *SearchPartitionRequest) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *SearchPartitionRequest) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *SearchPartitionRequest) SetClusters(v []string) {
	o.Clusters = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *SearchPartitionRequest) GetEncoding() RequestEncoding {
	if o == nil || IsNil(o.Encoding) {
		var ret RequestEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetEncodingOk() (*RequestEncoding, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *SearchPartitionRequest) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given RequestEncoding and assigns it to the Encoding field.
func (o *SearchPartitionRequest) SetEncoding(v RequestEncoding) {
	o.Encoding = &v
}

// GetEndTime returns the EndTime field value
func (o *SearchPartitionRequest) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SearchPartitionRequest) SetEndTime(v int64) {
	o.EndTime = v
}

// GetQueryFn returns the QueryFn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchPartitionRequest) GetQueryFn() string {
	if o == nil || IsNil(o.QueryFn.Get()) {
		var ret string
		return ret
	}
	return *o.QueryFn.Get()
}

// GetQueryFnOk returns a tuple with the QueryFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchPartitionRequest) GetQueryFnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryFn.Get(), o.QueryFn.IsSet()
}

// HasQueryFn returns a boolean if a field has been set.
func (o *SearchPartitionRequest) HasQueryFn() bool {
	if o != nil && o.QueryFn.IsSet() {
		return true
	}

	return false
}

// SetQueryFn gets a reference to the given NullableString and assigns it to the QueryFn field.
func (o *SearchPartitionRequest) SetQueryFn(v string) {
	o.QueryFn.Set(&v)
}

// SetQueryFnNil sets the value for QueryFn to be an explicit nil
func (o *SearchPartitionRequest) SetQueryFnNil() {
	o.QueryFn.Set(nil)
}

// UnsetQueryFn ensures that no value is present for QueryFn, not even an explicit nil
func (o *SearchPartitionRequest) UnsetQueryFn() {
	o.QueryFn.Unset()
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *SearchPartitionRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *SearchPartitionRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *SearchPartitionRequest) SetRegions(v []string) {
	o.Regions = v
}

// GetSql returns the Sql field value
func (o *SearchPartitionRequest) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *SearchPartitionRequest) SetSql(v string) {
	o.Sql = v
}

// GetStartTime returns the StartTime field value
func (o *SearchPartitionRequest) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SearchPartitionRequest) SetStartTime(v int64) {
	o.StartTime = v
}

// GetStreamingOutput returns the StreamingOutput field value if set, zero value otherwise.
func (o *SearchPartitionRequest) GetStreamingOutput() bool {
	if o == nil || IsNil(o.StreamingOutput) {
		var ret bool
		return ret
	}
	return *o.StreamingOutput
}

// GetStreamingOutputOk returns a tuple with the StreamingOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartitionRequest) GetStreamingOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.StreamingOutput) {
		return nil, false
	}
	return o.StreamingOutput, true
}

// HasStreamingOutput returns a boolean if a field has been set.
func (o *SearchPartitionRequest) HasStreamingOutput() bool {
	if o != nil && !IsNil(o.StreamingOutput) {
		return true
	}

	return false
}

// SetStreamingOutput gets a reference to the given bool and assigns it to the StreamingOutput field.
func (o *SearchPartitionRequest) SetStreamingOutput(v bool) {
	o.StreamingOutput = &v
}

func (o SearchPartitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPartitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	toSerialize["end_time"] = o.EndTime
	if o.QueryFn.IsSet() {
		toSerialize["query_fn"] = o.QueryFn.Get()
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	toSerialize["sql"] = o.Sql
	toSerialize["start_time"] = o.StartTime
	if !IsNil(o.StreamingOutput) {
		toSerialize["streaming_output"] = o.StreamingOutput
	}
	return toSerialize, nil
}

func (o *SearchPartitionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end_time",
		"sql",
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

	varSearchPartitionRequest := _SearchPartitionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchPartitionRequest)

	if err != nil {
		return err
	}

	*o = SearchPartitionRequest(varSearchPartitionRequest)

	return err
}

type NullableSearchPartitionRequest struct {
	value *SearchPartitionRequest
	isSet bool
}

func (v NullableSearchPartitionRequest) Get() *SearchPartitionRequest {
	return v.value
}

func (v *NullableSearchPartitionRequest) Set(val *SearchPartitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPartitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPartitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPartitionRequest(val *SearchPartitionRequest) *NullableSearchPartitionRequest {
	return &NullableSearchPartitionRequest{value: val, isSet: true}
}

func (v NullableSearchPartitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPartitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
