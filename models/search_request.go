package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequest{}

// SearchRequest struct for SearchRequest
type SearchRequest struct {
	Clusters           []string                   `json:"clusters,omitempty"`
	Encoding           *RequestEncoding           `json:"encoding,omitempty"`
	Query              SearchQuery                `json:"query"`
	Regions            []string                   `json:"regions,omitempty"`
	SearchEventContext NullableSearchEventContext `json:"search_event_context,omitempty"`
	SearchType         NullableSearchEventType    `json:"search_type,omitempty"`
	Timeout            *int64                     `json:"timeout,omitempty"`
	UseCache           NullableBool               `json:"use_cache,omitempty"`
}

// NewSearchRequest instantiates a new SearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequest(query SearchQuery) *SearchRequest {
	this := SearchRequest{}
	timeout := int64(5)
	this.Query = query
	this.Timeout = &timeout
	this.UseCache.Set(new(bool))
	*this.UseCache.Get() = true
	return &this
}

type _SearchRequest SearchRequest

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *SearchRequest) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *SearchRequest) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *SearchRequest) SetClusters(v []string) {
	o.Clusters = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *SearchRequest) GetEncoding() RequestEncoding {
	if o == nil || IsNil(o.Encoding) {
		var ret RequestEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetEncodingOk() (*RequestEncoding, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *SearchRequest) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given RequestEncoding and assigns it to the Encoding field.
func (o *SearchRequest) SetEncoding(v RequestEncoding) {
	o.Encoding = &v
}

// GetQuery returns the Query field value
func (o *SearchRequest) GetQuery() SearchQuery {
	if o == nil {
		var ret SearchQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetQueryOk() (*SearchQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchRequest) SetQuery(v SearchQuery) {
	o.Query = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *SearchRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *SearchRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *SearchRequest) SetRegions(v []string) {
	o.Regions = v
}

// GetSearchEventContext returns the SearchEventContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetSearchEventContext() SearchEventContext {
	if o == nil || IsNil(o.SearchEventContext.Get()) {
		var ret SearchEventContext
		return ret
	}
	return *o.SearchEventContext.Get()
}

// GetSearchEventContextOk returns a tuple with the SearchEventContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetSearchEventContextOk() (*SearchEventContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchEventContext.Get(), o.SearchEventContext.IsSet()
}

// HasSearchEventContext returns a boolean if a field has been set.
func (o *SearchRequest) HasSearchEventContext() bool {
	if o != nil && o.SearchEventContext.IsSet() {
		return true
	}

	return false
}

// SetSearchEventContext gets a reference to the given NullableSearchEventContext and assigns it to the SearchEventContext field.
func (o *SearchRequest) SetSearchEventContext(v SearchEventContext) {
	o.SearchEventContext.Set(&v)
}

// SetSearchEventContextNil sets the value for SearchEventContext to be an explicit nil
func (o *SearchRequest) SetSearchEventContextNil() {
	o.SearchEventContext.Set(nil)
}

// UnsetSearchEventContext ensures that no value is present for SearchEventContext, not even an explicit nil
func (o *SearchRequest) UnsetSearchEventContext() {
	o.SearchEventContext.Unset()
}

// GetSearchType returns the SearchType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetSearchType() SearchEventType {
	if o == nil || IsNil(o.SearchType.Get()) {
		var ret SearchEventType
		return ret
	}
	return *o.SearchType.Get()
}

// GetSearchTypeOk returns a tuple with the SearchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetSearchTypeOk() (*SearchEventType, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchType.Get(), o.SearchType.IsSet()
}

// HasSearchType returns a boolean if a field has been set.
func (o *SearchRequest) HasSearchType() bool {
	if o != nil && o.SearchType.IsSet() {
		return true
	}

	return false
}

// SetSearchType gets a reference to the given NullableSearchEventType and assigns it to the SearchType field.
func (o *SearchRequest) SetSearchType(v SearchEventType) {
	o.SearchType.Set(&v)
}

// SetSearchTypeNil sets the value for SearchType to be an explicit nil
func (o *SearchRequest) SetSearchTypeNil() {
	o.SearchType.Set(nil)
}

// UnsetSearchType ensures that no value is present for SearchType, not even an explicit nil
func (o *SearchRequest) UnsetSearchType() {
	o.SearchType.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SearchRequest) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SearchRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SearchRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUseCache returns the UseCache field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetUseCache() bool {
	if o == nil || IsNil(o.UseCache.Get()) {
		var ret bool
		return ret
	}
	return *o.UseCache.Get()
}

// GetUseCacheOk returns a tuple with the UseCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetUseCacheOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseCache.Get(), o.UseCache.IsSet()
}

// HasUseCache returns a boolean if a field has been set.
func (o *SearchRequest) HasUseCache() bool {
	if o != nil && o.UseCache.IsSet() {
		return true
	}

	return false
}

// SetUseCache gets a reference to the given NullableBool and assigns it to the UseCache field.
func (o *SearchRequest) SetUseCache(v bool) {
	o.UseCache.Set(&v)
}

// SetUseCacheNil sets the value for UseCache to be an explicit nil
func (o *SearchRequest) SetUseCacheNil() {
	o.UseCache.Set(nil)
}

// UnsetUseCache ensures that no value is present for UseCache, not even an explicit nil
func (o *SearchRequest) UnsetUseCache() {
	o.UseCache.Unset()
}

func (o SearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	toSerialize["query"] = o.Query
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if o.SearchEventContext.IsSet() {
		toSerialize["search_event_context"] = o.SearchEventContext.Get()
	}
	if o.SearchType.IsSet() {
		toSerialize["search_type"] = o.SearchType.Get()
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if o.UseCache.IsSet() {
		toSerialize["use_cache"] = o.UseCache.Get()
	}
	return toSerialize, nil
}

func (o *SearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varSearchRequest := _SearchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchRequest)

	if err != nil {
		return err
	}

	*o = SearchRequest(varSearchRequest)

	return err
}

type NullableSearchRequest struct {
	value *SearchRequest
	isSet bool
}

func (v NullableSearchRequest) Get() *SearchRequest {
	return v.value
}

func (v *NullableSearchRequest) Set(val *SearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequest(val *SearchRequest) *NullableSearchRequest {
	return &NullableSearchRequest{value: val, isSet: true}
}

func (v NullableSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
