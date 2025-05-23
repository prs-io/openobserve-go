package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PanelFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PanelFields{}

// PanelFields struct for PanelFields
type PanelFields struct {
	Filter     []PanelFilter `json:"filter"`
	Stream     string        `json:"stream"`
	StreamType StreamType    `json:"stream_type"`
	X          []AxisItem    `json:"x"`
	Y          []AxisItem    `json:"y"`
}

type _PanelFields PanelFields

// NewPanelFields instantiates a new PanelFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanelFields(filter []PanelFilter, stream string, streamType StreamType, x []AxisItem, y []AxisItem) *PanelFields {
	this := PanelFields{}
	this.Filter = filter
	this.Stream = stream
	this.StreamType = streamType
	this.X = x
	this.Y = y
	return &this
}

// GetFilter returns the Filter field value
func (o *PanelFields) GetFilter() []PanelFilter {
	if o == nil {
		var ret []PanelFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *PanelFields) GetFilterOk() ([]PanelFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *PanelFields) SetFilter(v []PanelFilter) {
	o.Filter = v
}

// GetStream returns the Stream field value
func (o *PanelFields) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *PanelFields) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *PanelFields) SetStream(v string) {
	o.Stream = v
}

// GetStreamType returns the StreamType field value
func (o *PanelFields) GetStreamType() StreamType {
	if o == nil {
		var ret StreamType
		return ret
	}

	return o.StreamType
}

// GetStreamTypeOk returns a tuple with the StreamType field value
// and a boolean to check if the value has been set.
func (o *PanelFields) GetStreamTypeOk() (*StreamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamType, true
}

// SetStreamType sets field value
func (o *PanelFields) SetStreamType(v StreamType) {
	o.StreamType = v
}

// GetX returns the X field value
func (o *PanelFields) GetX() []AxisItem {
	if o == nil {
		var ret []AxisItem
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *PanelFields) GetXOk() ([]AxisItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.X, true
}

// SetX sets field value
func (o *PanelFields) SetX(v []AxisItem) {
	o.X = v
}

// GetY returns the Y field value
func (o *PanelFields) GetY() []AxisItem {
	if o == nil {
		var ret []AxisItem
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *PanelFields) GetYOk() ([]AxisItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Y, true
}

// SetY sets field value
func (o *PanelFields) SetY(v []AxisItem) {
	o.Y = v
}

func (o PanelFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PanelFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	toSerialize["stream"] = o.Stream
	toSerialize["stream_type"] = o.StreamType
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *PanelFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter",
		"stream",
		"stream_type",
		"x",
		"y",
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

	varPanelFields := _PanelFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPanelFields)

	if err != nil {
		return err
	}

	*o = PanelFields(varPanelFields)

	return err
}

type NullablePanelFields struct {
	value *PanelFields
	isSet bool
}

func (v NullablePanelFields) Get() *PanelFields {
	return v.value
}

func (v *NullablePanelFields) Set(val *PanelFields) {
	v.value = val
	v.isSet = true
}

func (v NullablePanelFields) IsSet() bool {
	return v.isSet
}

func (v *NullablePanelFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanelFields(val *PanelFields) *NullablePanelFields {
	return &NullablePanelFields{value: val, isSet: true}
}

func (v NullablePanelFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanelFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
