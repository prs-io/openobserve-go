package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Variables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Variables{}

// Variables struct for Variables
type Variables struct {
	List []VariableList `json:"list"`
}

type _Variables Variables

// NewVariables instantiates a new Variables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariables(list []VariableList) *Variables {
	this := Variables{}
	this.List = list
	return &this
}

// GetList returns the List field value
func (o *Variables) GetList() []VariableList {
	if o == nil {
		var ret []VariableList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *Variables) GetListOk() ([]VariableList, bool) {
	if o == nil {
		return nil, false
	}
	return o.List, true
}

// SetList sets field value
func (o *Variables) SetList(v []VariableList) {
	o.List = v
}

func (o Variables) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Variables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

func (o *Variables) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list",
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

	varVariables := _Variables{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariables)

	if err != nil {
		return err
	}

	*o = Variables(varVariables)

	return err
}

type NullableVariables struct {
	value *Variables
	isSet bool
}

func (v NullableVariables) Get() *Variables {
	return v.value
}

func (v *NullableVariables) Set(val *Variables) {
	v.value = val
	v.isSet = true
}

func (v NullableVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariables(val *Variables) *NullableVariables {
	return &NullableVariables{value: val, isSet: true}
}

func (v NullableVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
