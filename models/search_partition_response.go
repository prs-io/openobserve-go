package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SearchPartitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPartitionResponse{}

// SearchPartitionResponse struct for SearchPartitionResponse
type SearchPartitionResponse struct {
	CompressedSize    int32          `json:"compressed_size"`
	FileNum           int32          `json:"file_num"`
	HistogramInterval NullableInt64  `json:"histogram_interval,omitempty"`
	MaxQueryRange     int64          `json:"max_query_range"`
	OrderBy           OrderBy        `json:"order_by"`
	OriginalSize      int64          `json:"original_size"`
	Partitions        [][]int64      `json:"partitions"`
	Records           int32          `json:"records"`
	StreamingAggs     bool           `json:"streaming_aggs"`
	StreamingId       NullableString `json:"streaming_id,omitempty"`
	StreamingOutput   bool           `json:"streaming_output"`
	TraceId           string         `json:"trace_id"`
	Limit             int32          `json:"limit"`
}

type _SearchPartitionResponse SearchPartitionResponse

// NewSearchPartitionResponse instantiates a new SearchPartitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPartitionResponse(compressedSize int32, fileNum int32, maxQueryRange int64, orderBy OrderBy, limit int32, originalSize int64, partitions [][]int64, records int32, streamingAggs bool, streamingOutput bool, traceId string) *SearchPartitionResponse {
	this := SearchPartitionResponse{}
	this.CompressedSize = compressedSize
	this.FileNum = fileNum
	this.MaxQueryRange = maxQueryRange
	this.OrderBy = orderBy
	this.Limit = limit
	this.OriginalSize = originalSize
	this.Partitions = partitions
	this.Records = records
	this.StreamingAggs = streamingAggs
	this.StreamingOutput = streamingOutput
	this.TraceId = traceId
	return &this
}

// GetCompressedSize returns the CompressedSize field value
func (o *SearchPartitionResponse) GetCompressedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompressedSize
}

// GetCompressedSizeOk returns a tuple with the CompressedSize field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetCompressedSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompressedSize, true
}

// SetCompressedSize sets field value
func (o *SearchPartitionResponse) SetCompressedSize(v int32) {
	o.CompressedSize = v
}

// GetFileNum returns the FileNum field value
func (o *SearchPartitionResponse) GetFileNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileNum
}

// GetFileNumOk returns a tuple with the FileNum field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetFileNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileNum, true
}

// SetFileNum sets field value
func (o *SearchPartitionResponse) SetFileNum(v int32) {
	o.FileNum = v
}

// GetHistogramInterval returns the HistogramInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchPartitionResponse) GetHistogramInterval() int64 {
	if o == nil || IsNil(o.HistogramInterval.Get()) {
		var ret int64
		return ret
	}
	return *o.HistogramInterval.Get()
}

// GetHistogramIntervalOk returns a tuple with the HistogramInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchPartitionResponse) GetHistogramIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HistogramInterval.Get(), o.HistogramInterval.IsSet()
}

// HasHistogramInterval returns a boolean if a field has been set.
func (o *SearchPartitionResponse) HasHistogramInterval() bool {
	if o != nil && o.HistogramInterval.IsSet() {
		return true
	}

	return false
}

// SetHistogramInterval gets a reference to the given NullableInt64 and assigns it to the HistogramInterval field.
func (o *SearchPartitionResponse) SetHistogramInterval(v int64) {
	o.HistogramInterval.Set(&v)
}

// SetHistogramIntervalNil sets the value for HistogramInterval to be an explicit nil
func (o *SearchPartitionResponse) SetHistogramIntervalNil() {
	o.HistogramInterval.Set(nil)
}

// UnsetHistogramInterval ensures that no value is present for HistogramInterval, not even an explicit nil
func (o *SearchPartitionResponse) UnsetHistogramInterval() {
	o.HistogramInterval.Unset()
}

// GetMaxQueryRange returns the MaxQueryRange field value
func (o *SearchPartitionResponse) GetMaxQueryRange() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxQueryRange
}

// GetMaxQueryRangeOk returns a tuple with the MaxQueryRange field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetMaxQueryRangeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxQueryRange, true
}

// SetMaxQueryRange sets field value
func (o *SearchPartitionResponse) SetMaxQueryRange(v int64) {
	o.MaxQueryRange = v
}

// GetOrderBy returns the OrderBy field value
func (o *SearchPartitionResponse) GetOrderBy() OrderBy {
	if o == nil {
		var ret OrderBy
		return ret
	}

	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetOrderByOk() (*OrderBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value
func (o *SearchPartitionResponse) SetOrderBy(v OrderBy) {
	o.OrderBy = v
}

// GetOriginalSize returns the OriginalSize field value
func (o *SearchPartitionResponse) GetOriginalSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetOriginalSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalSize, true
}

// SetOriginalSize sets field value
func (o *SearchPartitionResponse) SetOriginalSize(v int64) {
	o.OriginalSize = v
}

// GetPartitions returns the Partitions field value
func (o *SearchPartitionResponse) GetPartitions() [][]int64 {
	if o == nil {
		var ret [][]int64
		return ret
	}

	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetPartitionsOk() ([][]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Partitions, true
}

// SetPartitions sets field value
func (o *SearchPartitionResponse) SetPartitions(v [][]int64) {
	o.Partitions = v
}

// GetRecords returns the Records field value
func (o *SearchPartitionResponse) GetRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value
func (o *SearchPartitionResponse) SetRecords(v int32) {
	o.Records = v
}

// GetStreamingAggs returns the StreamingAggs field value
func (o *SearchPartitionResponse) GetStreamingAggs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StreamingAggs
}

// GetStreamingAggsOk returns a tuple with the StreamingAggs field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetStreamingAggsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingAggs, true
}

// SetStreamingAggs sets field value
func (o *SearchPartitionResponse) SetStreamingAggs(v bool) {
	o.StreamingAggs = v
}

// GetStreamingId returns the StreamingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchPartitionResponse) GetStreamingId() string {
	if o == nil || IsNil(o.StreamingId.Get()) {
		var ret string
		return ret
	}
	return *o.StreamingId.Get()
}

// GetStreamingIdOk returns a tuple with the StreamingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchPartitionResponse) GetStreamingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamingId.Get(), o.StreamingId.IsSet()
}

// HasStreamingId returns a boolean if a field has been set.
func (o *SearchPartitionResponse) HasStreamingId() bool {
	if o != nil && o.StreamingId.IsSet() {
		return true
	}

	return false
}

// SetStreamingId gets a reference to the given NullableString and assigns it to the StreamingId field.
func (o *SearchPartitionResponse) SetStreamingId(v string) {
	o.StreamingId.Set(&v)
}

// SetStreamingIdNil sets the value for StreamingId to be an explicit nil
func (o *SearchPartitionResponse) SetStreamingIdNil() {
	o.StreamingId.Set(nil)
}

// UnsetStreamingId ensures that no value is present for StreamingId, not even an explicit nil
func (o *SearchPartitionResponse) UnsetStreamingId() {
	o.StreamingId.Unset()
}

// GetStreamingOutput returns the StreamingOutput field value
func (o *SearchPartitionResponse) GetStreamingOutput() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StreamingOutput
}

// GetStreamingOutputOk returns a tuple with the StreamingOutput field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetStreamingOutputOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamingOutput, true
}

// SetStreamingOutput sets field value
func (o *SearchPartitionResponse) SetStreamingOutput(v bool) {
	o.StreamingOutput = v
}

// GetTraceId returns the TraceId field value
func (o *SearchPartitionResponse) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *SearchPartitionResponse) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *SearchPartitionResponse) SetTraceId(v string) {
	o.TraceId = v
}

func (o SearchPartitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPartitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["compressed_size"] = o.CompressedSize
	toSerialize["file_num"] = o.FileNum
	if o.HistogramInterval.IsSet() {
		toSerialize["histogram_interval"] = o.HistogramInterval.Get()
	}
	toSerialize["max_query_range"] = o.MaxQueryRange
	toSerialize["order_by"] = o.OrderBy
	toSerialize["original_size"] = o.OriginalSize
	toSerialize["partitions"] = o.Partitions
	toSerialize["records"] = o.Records
	toSerialize["streaming_aggs"] = o.StreamingAggs
	if o.StreamingId.IsSet() {
		toSerialize["streaming_id"] = o.StreamingId.Get()
	}
	toSerialize["streaming_output"] = o.StreamingOutput
	toSerialize["trace_id"] = o.TraceId
	return toSerialize, nil
}

func (o *SearchPartitionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"compressed_size",
		"file_num",
		"max_query_range",
		"order_by",
		"limit",
		"original_size",
		"partitions",
		"records",
		"streaming_aggs",
		"streaming_output",
		"trace_id",
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

	varSearchPartitionResponse := _SearchPartitionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchPartitionResponse)

	if err != nil {
		return err
	}

	*o = SearchPartitionResponse(varSearchPartitionResponse)

	return err
}

type NullableSearchPartitionResponse struct {
	value *SearchPartitionResponse
	isSet bool
}

func (v NullableSearchPartitionResponse) Get() *SearchPartitionResponse {
	return v.value
}

func (v *NullableSearchPartitionResponse) Set(val *SearchPartitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPartitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPartitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPartitionResponse(val *SearchPartitionResponse) *NullableSearchPartitionResponse {
	return &NullableSearchPartitionResponse{value: val, isSet: true}
}

func (v NullableSearchPartitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPartitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
