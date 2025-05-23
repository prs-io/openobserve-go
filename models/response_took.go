package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ResponseTook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTook{}

// ResponseTook struct for ResponseTook
type ResponseTook struct {
	SearchTook       int32              `json:"search_took"`
	CacheTook        int32              `json:"cache_took"`
	FileListTook     int32              `json:"file_list_took"`
	ClusterTotal     *int32             `json:"cluster_total,omitempty"`
	ClusterWaitQueue *int32             `json:"cluster_wait_queue,omitempty"`
	IdxTook          int32              `json:"idx_took"`
	Nodes            []ResponseNodeTook `json:"nodes,omitempty"`
	Total            int32              `json:"total"`
	WaitQueue        int32              `json:"wait_in_queue"`
}

type _ResponseTook ResponseTook

// NewResponseTook instantiates a new ResponseTook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTook(searchTook int32, cacheTook int32, fileListTook int32, idxTook int32, total int32, waitQueue int32) *ResponseTook {
	this := ResponseTook{}
	this.SearchTook = searchTook
	this.CacheTook = cacheTook
	this.FileListTook = fileListTook
	this.ClusterWaitQueue = nil
	this.IdxTook = idxTook
	this.Total = total
	this.WaitQueue = waitQueue
	return &this
}

// GetClusterTotal returns the ClusterTotal field value
func (o *ResponseTook) GetClusterTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.ClusterTotal
}

// GetClusterTotalOk returns a tuple with the ClusterTotal field value
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetClusterTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterTotal, true
}

// SetClusterTotal sets field value
func (o *ResponseTook) SetClusterTotal(v int32) {
	o.ClusterTotal = &v
}

// GetClusterWaitQueue returns the ClusterWaitQueue field value
func (o *ResponseTook) GetClusterWaitQueue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.ClusterWaitQueue
}

// GetClusterWaitQueueOk returns a tuple with the ClusterWaitQueue field value
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetClusterWaitQueueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterWaitQueue, true
}

// SetClusterWaitQueue sets field value
func (o *ResponseTook) SetClusterWaitQueue(v int32) {
	o.ClusterWaitQueue = &v
}

// GetIdxTook returns the IdxTook field value
func (o *ResponseTook) GetIdxTook() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdxTook
}

// GetIdxTookOk returns a tuple with the IdxTook field value
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetIdxTookOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdxTook, true
}

// SetIdxTook sets field value
func (o *ResponseTook) SetIdxTook(v int32) {
	o.IdxTook = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ResponseTook) GetNodes() []ResponseNodeTook {
	if o == nil || IsNil(o.Nodes) {
		var ret []ResponseNodeTook
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetNodesOk() ([]ResponseNodeTook, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ResponseTook) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []ResponseNodeTook and assigns it to the Nodes field.
func (o *ResponseTook) SetNodes(v []ResponseNodeTook) {
	o.Nodes = v
}

// GetTotal returns the Total field value
func (o *ResponseTook) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseTook) SetTotal(v int32) {
	o.Total = v
}

// GetWaitQueue returns the WaitQueue field value
func (o *ResponseTook) GetWaitQueue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WaitQueue
}

// GetWaitQueueOk returns a tuple with the WaitQueue field value
// and a boolean to check if the value has been set.
func (o *ResponseTook) GetWaitQueueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitQueue, true
}

// SetWaitQueue sets field value
func (o *ResponseTook) SetWaitQueue(v int32) {
	o.WaitQueue = v
}

func (o ResponseTook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_total"] = o.ClusterTotal
	toSerialize["cluster_wait_queue"] = o.ClusterWaitQueue
	toSerialize["idx_took"] = o.IdxTook
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	toSerialize["total"] = o.Total
	toSerialize["wait_queue"] = o.WaitQueue
	return toSerialize, nil
}

func (o *ResponseTook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		// "cluster_total",
		// "cluster_wait_queue",
		"cache_took",
		"file_list_took",
		"idx_took",
		"total",
		"wait_in_queue",
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

	varResponseTook := _ResponseTook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseTook)

	if err != nil {
		return err
	}

	*o = ResponseTook(varResponseTook)

	return err
}

type NullableResponseTook struct {
	value *ResponseTook
	isSet bool
}

func (v NullableResponseTook) Get() *ResponseTook {
	return v.value
}

func (v *NullableResponseTook) Set(val *ResponseTook) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTook) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTook(val *ResponseTook) *NullableResponseTook {
	return &NullableResponseTook{value: val, isSet: true}
}

func (v NullableResponseTook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
