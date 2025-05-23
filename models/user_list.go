package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserList{}

// UserList struct for UserList
type UserList struct {
	Data []UserResponse `json:"data"`
}

type _UserList UserList

// NewUserList instantiates a new UserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserList(data []UserResponse) *UserList {
	this := UserList{}
	this.Data = data
	return &this
}

// GetData returns the Data field value
func (o *UserList) GetData() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserList) GetDataOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UserList) SetData(v []UserResponse) {
	o.Data = v
}

func (o UserList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UserList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varUserList := _UserList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserList)

	if err != nil {
		return err
	}

	*o = UserList(varUserList)

	return err
}

type NullableUserList struct {
	value *UserList
	isSet bool
}

func (v NullableUserList) Get() *UserList {
	return v.value
}

func (v *NullableUserList) Set(val *UserList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserList(val *UserList) *NullableUserList {
	return &NullableUserList{value: val, isSet: true}
}

func (v NullableUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
