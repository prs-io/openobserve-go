package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CustomFieldsOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldsOption{}

// CustomFieldsOption struct for CustomFieldsOption
type CustomFieldsOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type _CustomFieldsOption CustomFieldsOption

// NewCustomFieldsOption instantiates a new CustomFieldsOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldsOption(label string, value string) *CustomFieldsOption {
	this := CustomFieldsOption{}
	this.Label = label
	this.Value = value
	return &this
}

// GetLabel returns the Label field value
func (o *CustomFieldsOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CustomFieldsOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CustomFieldsOption) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *CustomFieldsOption) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CustomFieldsOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CustomFieldsOption) SetValue(v string) {
	o.Value = v
}

func (o CustomFieldsOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldsOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CustomFieldsOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"value",
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

	varCustomFieldsOption := _CustomFieldsOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomFieldsOption)

	if err != nil {
		return err
	}

	*o = CustomFieldsOption(varCustomFieldsOption)

	return err
}

type NullableCustomFieldsOption struct {
	value *CustomFieldsOption
	isSet bool
}

func (v NullableCustomFieldsOption) Get() *CustomFieldsOption {
	return v.value
}

func (v *NullableCustomFieldsOption) Set(val *CustomFieldsOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldsOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldsOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldsOption(val *CustomFieldsOption) *NullableCustomFieldsOption {
	return &NullableCustomFieldsOption{value: val, isSet: true}
}

func (v NullableCustomFieldsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldsOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
