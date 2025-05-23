package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VariableList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableList{}

// VariableList struct for VariableList
type VariableList struct {
	Label     string               `json:"label"`
	Name      string               `json:"name"`
	Options   []CustomFieldsOption `json:"options,omitempty"`
	QueryData NullableQueryData    `json:"query_data,omitempty"`
	Type      string               `json:"type"`
	Value     NullableString       `json:"value,omitempty"`
}

type _VariableList VariableList

// NewVariableList instantiates a new VariableList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableList(label string, name string, type_ string) *VariableList {
	this := VariableList{}
	this.Label = label
	this.Name = name
	this.Type = type_
	return &this
}

// GetLabel returns the Label field value
func (o *VariableList) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VariableList) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VariableList) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *VariableList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableList) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableList) GetOptions() []CustomFieldsOption {
	if o == nil {
		var ret []CustomFieldsOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableList) GetOptionsOk() ([]CustomFieldsOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *VariableList) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomFieldsOption and assigns it to the Options field.
func (o *VariableList) SetOptions(v []CustomFieldsOption) {
	o.Options = v
}

// GetQueryData returns the QueryData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableList) GetQueryData() QueryData {
	if o == nil || IsNil(o.QueryData.Get()) {
		var ret QueryData
		return ret
	}
	return *o.QueryData.Get()
}

// GetQueryDataOk returns a tuple with the QueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableList) GetQueryDataOk() (*QueryData, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryData.Get(), o.QueryData.IsSet()
}

// HasQueryData returns a boolean if a field has been set.
func (o *VariableList) HasQueryData() bool {
	if o != nil && o.QueryData.IsSet() {
		return true
	}

	return false
}

// SetQueryData gets a reference to the given NullableQueryData and assigns it to the QueryData field.
func (o *VariableList) SetQueryData(v QueryData) {
	o.QueryData.Set(&v)
}

// SetQueryDataNil sets the value for QueryData to be an explicit nil
func (o *VariableList) SetQueryDataNil() {
	o.QueryData.Set(nil)
}

// UnsetQueryData ensures that no value is present for QueryData, not even an explicit nil
func (o *VariableList) UnsetQueryData() {
	o.QueryData.Unset()
}

// GetType returns the Type field value
func (o *VariableList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VariableList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VariableList) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariableList) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariableList) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *VariableList) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *VariableList) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *VariableList) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *VariableList) UnsetValue() {
	o.Value.Unset()
}

func (o VariableList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.QueryData.IsSet() {
		toSerialize["query_data"] = o.QueryData.Get()
	}
	toSerialize["type"] = o.Type
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

func (o *VariableList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"type",
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

	varVariableList := _VariableList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableList)

	if err != nil {
		return err
	}

	*o = VariableList(varVariableList)

	return err
}

type NullableVariableList struct {
	value *VariableList
	isSet bool
}

func (v NullableVariableList) Get() *VariableList {
	return v.value
}

func (v *NullableVariableList) Set(val *VariableList) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableList) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableList(val *VariableList) *NullableVariableList {
	return &NullableVariableList{value: val, isSet: true}
}

func (v NullableVariableList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
