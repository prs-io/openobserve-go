package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the QueryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryData{}

// QueryData struct for QueryData
type QueryData struct {
	Field         string        `json:"field"`
	MaxRecordSize NullableInt64 `json:"max_record_size,omitempty"`
	Stream        string        `json:"stream"`
	StreamType    StreamType    `json:"stream_type"`
}

type _QueryData QueryData

// NewQueryData instantiates a new QueryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryData(field string, stream string, streamType StreamType) *QueryData {
	this := QueryData{}
	this.Field = field
	this.Stream = stream
	this.StreamType = streamType
	return &this
}

// GetField returns the Field field value
func (o *QueryData) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *QueryData) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *QueryData) SetField(v string) {
	o.Field = v
}

// GetMaxRecordSize returns the MaxRecordSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryData) GetMaxRecordSize() int64 {
	if o == nil || IsNil(o.MaxRecordSize.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRecordSize.Get()
}

// GetMaxRecordSizeOk returns a tuple with the MaxRecordSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryData) GetMaxRecordSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRecordSize.Get(), o.MaxRecordSize.IsSet()
}

// HasMaxRecordSize returns a boolean if a field has been set.
func (o *QueryData) HasMaxRecordSize() bool {
	if o != nil && o.MaxRecordSize.IsSet() {
		return true
	}

	return false
}

// SetMaxRecordSize gets a reference to the given NullableInt64 and assigns it to the MaxRecordSize field.
func (o *QueryData) SetMaxRecordSize(v int64) {
	o.MaxRecordSize.Set(&v)
}

// SetMaxRecordSizeNil sets the value for MaxRecordSize to be an explicit nil
func (o *QueryData) SetMaxRecordSizeNil() {
	o.MaxRecordSize.Set(nil)
}

// UnsetMaxRecordSize ensures that no value is present for MaxRecordSize, not even an explicit nil
func (o *QueryData) UnsetMaxRecordSize() {
	o.MaxRecordSize.Unset()
}

// GetStream returns the Stream field value
func (o *QueryData) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *QueryData) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *QueryData) SetStream(v string) {
	o.Stream = v
}

// GetStreamType returns the StreamType field value
func (o *QueryData) GetStreamType() StreamType {
	if o == nil {
		var ret StreamType
		return ret
	}

	return o.StreamType
}

// GetStreamTypeOk returns a tuple with the StreamType field value
// and a boolean to check if the value has been set.
func (o *QueryData) GetStreamTypeOk() (*StreamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamType, true
}

// SetStreamType sets field value
func (o *QueryData) SetStreamType(v StreamType) {
	o.StreamType = v
}

func (o QueryData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	if o.MaxRecordSize.IsSet() {
		toSerialize["max_record_size"] = o.MaxRecordSize.Get()
	}
	toSerialize["stream"] = o.Stream
	toSerialize["stream_type"] = o.StreamType
	return toSerialize, nil
}

func (o *QueryData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"stream",
		"stream_type",
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

	varQueryData := _QueryData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryData)

	if err != nil {
		return err
	}

	*o = QueryData(varQueryData)

	return err
}

type NullableQueryData struct {
	value *QueryData
	isSet bool
}

func (v NullableQueryData) Get() *QueryData {
	return v.value
}

func (v *NullableQueryData) Set(val *QueryData) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryData) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryData(val *QueryData) *NullableQueryData {
	return &NullableQueryData{value: val, isSet: true}
}

func (v NullableQueryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
