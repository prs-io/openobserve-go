package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserOrgRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserOrgRole{}

// UserOrgRole struct for UserOrgRole
type UserOrgRole struct {
	Role UserRole `json:"role"`
}

type _UserOrgRole UserOrgRole

// NewUserOrgRole instantiates a new UserOrgRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOrgRole(role UserRole) *UserOrgRole {
	this := UserOrgRole{}
	this.Role = role
	return &this
}

// GetRole returns the Role field value
func (o *UserOrgRole) GetRole() UserRole {
	if o == nil {
		var ret UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserOrgRole) GetRoleOk() (*UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserOrgRole) SetRole(v UserRole) {
	o.Role = v
}

func (o UserOrgRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserOrgRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *UserOrgRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
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

	varUserOrgRole := _UserOrgRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserOrgRole)

	if err != nil {
		return err
	}

	*o = UserOrgRole(varUserOrgRole)

	return err
}

type NullableUserOrgRole struct {
	value *UserOrgRole
	isSet bool
}

func (v NullableUserOrgRole) Get() *UserOrgRole {
	return v.value
}

func (v *NullableUserOrgRole) Set(val *UserOrgRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOrgRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOrgRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOrgRole(val *UserOrgRole) *NullableUserOrgRole {
	return &NullableUserOrgRole{value: val, isSet: true}
}

func (v NullableUserOrgRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOrgRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
