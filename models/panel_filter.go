package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PanelFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PanelFilter{}

// PanelFilter struct for PanelFilter
type PanelFilter struct {
	Column   string         `json:"column"`
	Operator NullableString `json:"operator,omitempty"`
	Type     string         `json:"type"`
	Value    NullableString `json:"value,omitempty"`
	Values   []string       `json:"values"`
}

type _PanelFilter PanelFilter

// NewPanelFilter instantiates a new PanelFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanelFilter(column string, type_ string, values []string) *PanelFilter {
	this := PanelFilter{}
	this.Column = column
	this.Type = type_
	this.Values = values
	return &this
}

// GetColumn returns the Column field value
func (o *PanelFilter) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *PanelFilter) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *PanelFilter) SetColumn(v string) {
	o.Column = v
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelFilter) GetOperator() string {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *PanelFilter) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *PanelFilter) SetOperator(v string) {
	o.Operator.Set(&v)
}

// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *PanelFilter) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *PanelFilter) UnsetOperator() {
	o.Operator.Unset()
}

// GetType returns the Type field value
func (o *PanelFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PanelFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PanelFilter) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelFilter) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *PanelFilter) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *PanelFilter) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *PanelFilter) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *PanelFilter) UnsetValue() {
	o.Value.Unset()
}

// GetValues returns the Values field value
func (o *PanelFilter) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *PanelFilter) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *PanelFilter) SetValues(v []string) {
	o.Values = v
}

func (o PanelFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PanelFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["column"] = o.Column
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	toSerialize["type"] = o.Type
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *PanelFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"column",
		"type",
		"values",
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

	varPanelFilter := _PanelFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPanelFilter)

	if err != nil {
		return err
	}

	*o = PanelFilter(varPanelFilter)

	return err
}

type NullablePanelFilter struct {
	value *PanelFilter
	isSet bool
}

func (v NullablePanelFilter) Get() *PanelFilter {
	return v.value
}

func (v *NullablePanelFilter) Set(val *PanelFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePanelFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePanelFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanelFilter(val *PanelFilter) *NullablePanelFilter {
	return &NullablePanelFilter{value: val, isSet: true}
}

func (v NullablePanelFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanelFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
