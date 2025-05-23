package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponse{}

// SearchResponse struct for SearchResponse
type SearchResponse struct {
	CachedRatio       int32                    `json:"cached_ratio"`
	Columns           []string                 `json:"columns,omitempty"`
	From              int64                    `json:"from"`
	FunctionError     *string                  `json:"function_error,omitempty"`
	HistogramInterval NullableInt64            `json:"histogram_interval,omitempty"`
	Hits              []map[string]interface{} `json:"hits"`
	IdxScanSize       int32                    `json:"idx_scan_size"`
	IsPartial         *bool                    `json:"is_partial,omitempty"`
	NewEndTime        NullableInt64            `json:"new_end_time,omitempty"`
	NewStartTime      NullableInt64            `json:"new_start_time,omitempty"`
	OrderBy           NullableOrderBy          `json:"order_by,omitempty"`
	ResponseType      *string                  `json:"response_type,omitempty"`
	ResultCacheRatio  *int32                   `json:"result_cache_ratio,omitempty"`
	ScanRecords       int32                    `json:"scan_records"`
	ScanSize          int32                    `json:"scan_size"`
	Size              int64                    `json:"size"`
	Took              int32                    `json:"took"`
	TookDetail        *NullableResponseTook    `json:"took_detail,omitempty"`
	Total             int32                    `json:"total"`
	TraceId           *string                  `json:"trace_id,omitempty"`
	WorkGroup         NullableString           `json:"work_group,omitempty"`
}

type _SearchResponse SearchResponse

// NewSearchResponse instantiates a new SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponse(cachedRatio int32, from int64, hits []map[string]interface{}, idxScanSize int32, scanRecords int32, scanSize int32, size int64, took int32, total int32) *SearchResponse {
	this := SearchResponse{}
	this.CachedRatio = cachedRatio
	this.From = from
	this.Hits = hits
	this.IdxScanSize = idxScanSize
	this.ScanRecords = scanRecords
	this.ScanSize = scanSize
	this.Size = size
	this.Took = took
	this.Total = total
	return &this
}

// GetCachedRatio returns the CachedRatio field value
func (o *SearchResponse) GetCachedRatio() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CachedRatio
}

// GetCachedRatioOk returns a tuple with the CachedRatio field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetCachedRatioOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CachedRatio, true
}

// SetCachedRatio sets field value
func (o *SearchResponse) SetCachedRatio(v int32) {
	o.CachedRatio = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SearchResponse) GetColumns() []string {
	if o == nil || IsNil(o.Columns) {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SearchResponse) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *SearchResponse) SetColumns(v []string) {
	o.Columns = v
}

// GetFrom returns the From field value
func (o *SearchResponse) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *SearchResponse) SetFrom(v int64) {
	o.From = v
}

// GetFunctionError returns the FunctionError field value if set, zero value otherwise.
func (o *SearchResponse) GetFunctionError() string {
	if o == nil || IsNil(o.FunctionError) {
		var ret string
		return ret
	}
	return *o.FunctionError
}

// GetFunctionErrorOk returns a tuple with the FunctionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetFunctionErrorOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionError) {
		return nil, false
	}
	return o.FunctionError, true
}

// HasFunctionError returns a boolean if a field has been set.
func (o *SearchResponse) HasFunctionError() bool {
	if o != nil && !IsNil(o.FunctionError) {
		return true
	}

	return false
}

// SetFunctionError gets a reference to the given string and assigns it to the FunctionError field.
func (o *SearchResponse) SetFunctionError(v string) {
	o.FunctionError = &v
}

// GetHistogramInterval returns the HistogramInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetHistogramInterval() int64 {
	if o == nil || IsNil(o.HistogramInterval.Get()) {
		var ret int64
		return ret
	}
	return *o.HistogramInterval.Get()
}

// GetHistogramIntervalOk returns a tuple with the HistogramInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetHistogramIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HistogramInterval.Get(), o.HistogramInterval.IsSet()
}

// HasHistogramInterval returns a boolean if a field has been set.
func (o *SearchResponse) HasHistogramInterval() bool {
	if o != nil && o.HistogramInterval.IsSet() {
		return true
	}

	return false
}

// SetHistogramInterval gets a reference to the given NullableInt64 and assigns it to the HistogramInterval field.
func (o *SearchResponse) SetHistogramInterval(v int64) {
	o.HistogramInterval.Set(&v)
}

// SetHistogramIntervalNil sets the value for HistogramInterval to be an explicit nil
func (o *SearchResponse) SetHistogramIntervalNil() {
	o.HistogramInterval.Set(nil)
}

// UnsetHistogramInterval ensures that no value is present for HistogramInterval, not even an explicit nil
func (o *SearchResponse) UnsetHistogramInterval() {
	o.HistogramInterval.Unset()
}

// GetHits returns the Hits field value
func (o *SearchResponse) GetHits() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetHitsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hits, true
}

// SetHits sets field value
func (o *SearchResponse) SetHits(v []map[string]interface{}) {
	o.Hits = v
}

// GetIdxScanSize returns the IdxScanSize field value
func (o *SearchResponse) GetIdxScanSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdxScanSize
}

// GetIdxScanSizeOk returns a tuple with the IdxScanSize field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetIdxScanSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdxScanSize, true
}

// SetIdxScanSize sets field value
func (o *SearchResponse) SetIdxScanSize(v int32) {
	o.IdxScanSize = v
}

// GetIsPartial returns the IsPartial field value if set, zero value otherwise.
func (o *SearchResponse) GetIsPartial() bool {
	if o == nil || IsNil(o.IsPartial) {
		var ret bool
		return ret
	}
	return *o.IsPartial
}

// GetIsPartialOk returns a tuple with the IsPartial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetIsPartialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartial) {
		return nil, false
	}
	return o.IsPartial, true
}

// HasIsPartial returns a boolean if a field has been set.
func (o *SearchResponse) HasIsPartial() bool {
	if o != nil && !IsNil(o.IsPartial) {
		return true
	}

	return false
}

// SetIsPartial gets a reference to the given bool and assigns it to the IsPartial field.
func (o *SearchResponse) SetIsPartial(v bool) {
	o.IsPartial = &v
}

// GetNewEndTime returns the NewEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetNewEndTime() int64 {
	if o == nil || IsNil(o.NewEndTime.Get()) {
		var ret int64
		return ret
	}
	return *o.NewEndTime.Get()
}

// GetNewEndTimeOk returns a tuple with the NewEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetNewEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewEndTime.Get(), o.NewEndTime.IsSet()
}

// HasNewEndTime returns a boolean if a field has been set.
func (o *SearchResponse) HasNewEndTime() bool {
	if o != nil && o.NewEndTime.IsSet() {
		return true
	}

	return false
}

// SetNewEndTime gets a reference to the given NullableInt64 and assigns it to the NewEndTime field.
func (o *SearchResponse) SetNewEndTime(v int64) {
	o.NewEndTime.Set(&v)
}

// SetNewEndTimeNil sets the value for NewEndTime to be an explicit nil
func (o *SearchResponse) SetNewEndTimeNil() {
	o.NewEndTime.Set(nil)
}

// UnsetNewEndTime ensures that no value is present for NewEndTime, not even an explicit nil
func (o *SearchResponse) UnsetNewEndTime() {
	o.NewEndTime.Unset()
}

// GetNewStartTime returns the NewStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetNewStartTime() int64 {
	if o == nil || IsNil(o.NewStartTime.Get()) {
		var ret int64
		return ret
	}
	return *o.NewStartTime.Get()
}

// GetNewStartTimeOk returns a tuple with the NewStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetNewStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewStartTime.Get(), o.NewStartTime.IsSet()
}

// HasNewStartTime returns a boolean if a field has been set.
func (o *SearchResponse) HasNewStartTime() bool {
	if o != nil && o.NewStartTime.IsSet() {
		return true
	}

	return false
}

// SetNewStartTime gets a reference to the given NullableInt64 and assigns it to the NewStartTime field.
func (o *SearchResponse) SetNewStartTime(v int64) {
	o.NewStartTime.Set(&v)
}

// SetNewStartTimeNil sets the value for NewStartTime to be an explicit nil
func (o *SearchResponse) SetNewStartTimeNil() {
	o.NewStartTime.Set(nil)
}

// UnsetNewStartTime ensures that no value is present for NewStartTime, not even an explicit nil
func (o *SearchResponse) UnsetNewStartTime() {
	o.NewStartTime.Unset()
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetOrderBy() OrderBy {
	if o == nil || IsNil(o.OrderBy.Get()) {
		var ret OrderBy
		return ret
	}
	return *o.OrderBy.Get()
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetOrderByOk() (*OrderBy, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderBy.Get(), o.OrderBy.IsSet()
}

// HasOrderBy returns a boolean if a field has been set.
func (o *SearchResponse) HasOrderBy() bool {
	if o != nil && o.OrderBy.IsSet() {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given NullableOrderBy and assigns it to the OrderBy field.
func (o *SearchResponse) SetOrderBy(v OrderBy) {
	o.OrderBy.Set(&v)
}

// SetOrderByNil sets the value for OrderBy to be an explicit nil
func (o *SearchResponse) SetOrderByNil() {
	o.OrderBy.Set(nil)
}

// UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
func (o *SearchResponse) UnsetOrderBy() {
	o.OrderBy.Unset()
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *SearchResponse) GetResponseType() string {
	if o == nil || IsNil(o.ResponseType) {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResponseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *SearchResponse) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *SearchResponse) SetResponseType(v string) {
	o.ResponseType = &v
}

// GetResultCacheRatio returns the ResultCacheRatio field value if set, zero value otherwise.
func (o *SearchResponse) GetResultCacheRatio() int32 {
	if o == nil || IsNil(o.ResultCacheRatio) {
		var ret int32
		return ret
	}
	return *o.ResultCacheRatio
}

// GetResultCacheRatioOk returns a tuple with the ResultCacheRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultCacheRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.ResultCacheRatio) {
		return nil, false
	}
	return o.ResultCacheRatio, true
}

// HasResultCacheRatio returns a boolean if a field has been set.
func (o *SearchResponse) HasResultCacheRatio() bool {
	if o != nil && !IsNil(o.ResultCacheRatio) {
		return true
	}

	return false
}

// SetResultCacheRatio gets a reference to the given int32 and assigns it to the ResultCacheRatio field.
func (o *SearchResponse) SetResultCacheRatio(v int32) {
	o.ResultCacheRatio = &v
}

// GetScanRecords returns the ScanRecords field value
func (o *SearchResponse) GetScanRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScanRecords
}

// GetScanRecordsOk returns a tuple with the ScanRecords field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetScanRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanRecords, true
}

// SetScanRecords sets field value
func (o *SearchResponse) SetScanRecords(v int32) {
	o.ScanRecords = v
}

// GetScanSize returns the ScanSize field value
func (o *SearchResponse) GetScanSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScanSize
}

// GetScanSizeOk returns a tuple with the ScanSize field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetScanSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanSize, true
}

// SetScanSize sets field value
func (o *SearchResponse) SetScanSize(v int32) {
	o.ScanSize = v
}

// GetSize returns the Size field value
func (o *SearchResponse) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *SearchResponse) SetSize(v int64) {
	o.Size = v
}

// GetTook returns the Took field value
func (o *SearchResponse) GetTook() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Took
}

// GetTookOk returns a tuple with the Took field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTookOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Took, true
}

// SetTook sets field value
func (o *SearchResponse) SetTook(v int32) {
	o.Took = v
}

// GetTookDetail returns the TookDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetTookDetail() ResponseTook {
	if o == nil || IsNil(o.TookDetail.Get()) {
		var ret ResponseTook
		return ret
	}
	return *o.TookDetail.Get()
}

// GetTookDetailOk returns a tuple with the TookDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetTookDetailOk() (*ResponseTook, bool) {
	if o == nil {
		return nil, false
	}
	return o.TookDetail.Get(), o.TookDetail.IsSet()
}

// HasTookDetail returns a boolean if a field has been set.
func (o *SearchResponse) HasTookDetail() bool {
	if o != nil && o.TookDetail.IsSet() {
		return true
	}

	return false
}

// SetTookDetail gets a reference to the given NullableResponseTook and assigns it to the TookDetail field.
func (o *SearchResponse) SetTookDetail(v ResponseTook) {
	o.TookDetail.Set(&v)
}

// SetTookDetailNil sets the value for TookDetail to be an explicit nil
func (o *SearchResponse) SetTookDetailNil() {
	o.TookDetail.Set(nil)
}

// UnsetTookDetail ensures that no value is present for TookDetail, not even an explicit nil
func (o *SearchResponse) UnsetTookDetail() {
	o.TookDetail.Unset()
}

// GetTotal returns the Total field value
func (o *SearchResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SearchResponse) SetTotal(v int32) {
	o.Total = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SearchResponse) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SearchResponse) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SearchResponse) SetTraceId(v string) {
	o.TraceId = &v
}

// GetWorkGroup returns the WorkGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResponse) GetWorkGroup() string {
	if o == nil || IsNil(o.WorkGroup.Get()) {
		var ret string
		return ret
	}
	return *o.WorkGroup.Get()
}

// GetWorkGroupOk returns a tuple with the WorkGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResponse) GetWorkGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkGroup.Get(), o.WorkGroup.IsSet()
}

// HasWorkGroup returns a boolean if a field has been set.
func (o *SearchResponse) HasWorkGroup() bool {
	if o != nil && o.WorkGroup.IsSet() {
		return true
	}

	return false
}

// SetWorkGroup gets a reference to the given NullableString and assigns it to the WorkGroup field.
func (o *SearchResponse) SetWorkGroup(v string) {
	o.WorkGroup.Set(&v)
}

// SetWorkGroupNil sets the value for WorkGroup to be an explicit nil
func (o *SearchResponse) SetWorkGroupNil() {
	o.WorkGroup.Set(nil)
}

// UnsetWorkGroup ensures that no value is present for WorkGroup, not even an explicit nil
func (o *SearchResponse) UnsetWorkGroup() {
	o.WorkGroup.Unset()
}

func (o SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cached_ratio"] = o.CachedRatio
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["from"] = o.From
	if !IsNil(o.FunctionError) {
		toSerialize["function_error"] = o.FunctionError
	}
	if o.HistogramInterval.IsSet() {
		toSerialize["histogram_interval"] = o.HistogramInterval.Get()
	}
	toSerialize["hits"] = o.Hits
	toSerialize["idx_scan_size"] = o.IdxScanSize
	if !IsNil(o.IsPartial) {
		toSerialize["is_partial"] = o.IsPartial
	}
	if o.NewEndTime.IsSet() {
		toSerialize["new_end_time"] = o.NewEndTime.Get()
	}
	if o.NewStartTime.IsSet() {
		toSerialize["new_start_time"] = o.NewStartTime.Get()
	}
	if o.OrderBy.IsSet() {
		toSerialize["order_by"] = o.OrderBy.Get()
	}
	if !IsNil(o.ResponseType) {
		toSerialize["response_type"] = o.ResponseType
	}
	if !IsNil(o.ResultCacheRatio) {
		toSerialize["result_cache_ratio"] = o.ResultCacheRatio
	}
	toSerialize["scan_records"] = o.ScanRecords
	toSerialize["scan_size"] = o.ScanSize
	toSerialize["size"] = o.Size
	toSerialize["took"] = o.Took
	if o.TookDetail.IsSet() {
		toSerialize["took_detail"] = o.TookDetail.Get()
	}
	toSerialize["total"] = o.Total
	if !IsNil(o.TraceId) {
		toSerialize["trace_id"] = o.TraceId
	}
	if o.WorkGroup.IsSet() {
		toSerialize["work_group"] = o.WorkGroup.Get()
	}
	return toSerialize, nil
}

func (o *SearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cached_ratio",
		"from",
		"hits",
		"idx_scan_size",
		"scan_records",
		"scan_size",
		"size",
		"took",
		"total",
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

	varSearchResponse := _SearchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchResponse)

	if err != nil {
		return err
	}

	*o = SearchResponse(varSearchResponse)

	return err
}

type NullableSearchResponse struct {
	value *SearchResponse
	isSet bool
}

func (v NullableSearchResponse) Get() *SearchResponse {
	return v.value
}

func (v *NullableSearchResponse) Set(val *SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponse(val *SearchResponse) *NullableSearchResponse {
	return &NullableSearchResponse{value: val, isSet: true}
}

func (v NullableSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
