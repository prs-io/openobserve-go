package models

import (
	"encoding/json"
	"fmt"
)

// UserRole the model 'UserRole'
type UserRole string

// List of UserRole
const (
	ADMIN           UserRole = "admin"
	MEMBER          UserRole = "member"
	ROOT            UserRole = "root"
	VIEWER          UserRole = "viewer"
	USER            UserRole = "user"
	EDITOR          UserRole = "editor"
	SERVICE_ACCOUNT UserRole = "service_account"
)

// All allowed values of UserRole enum
var AllowedUserRoleEnumValues = []UserRole{
	"admin",
	"member",
	"root",
	"viewer",
	"user",
	"editor",
	"service_account",
}

func (v *UserRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserRole(value)
	for _, existing := range AllowedUserRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserRole", value)
}

// NewUserRoleFromValue returns a pointer to a valid UserRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserRoleFromValue(v string) (*UserRole, error) {
	ev := UserRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserRole: valid values are %v", v, AllowedUserRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserRole) IsValid() bool {
	for _, existing := range AllowedUserRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserRole value
func (v UserRole) Ptr() *UserRole {
	return &v
}

type NullableUserRole struct {
	value *UserRole
	isSet bool
}

func (v NullableUserRole) Get() *UserRole {
	return v.value
}

func (v *NullableUserRole) Set(val *UserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRole(val *UserRole) *NullableUserRole {
	return &NullableUserRole{value: val, isSet: true}
}

func (v NullableUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
