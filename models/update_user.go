package models

import (
	"encoding/json"
)

// checks if the UpdateUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUser{}

// UpdateUser struct for UpdateUser
type UpdateUser struct {
	ChangePassword *bool            `json:"change_password,omitempty"`
	FirstName      NullableString   `json:"first_name,omitempty"`
	LastName       NullableString   `json:"last_name,omitempty"`
	NewPassword    NullableString   `json:"new_password,omitempty"`
	OldPassword    NullableString   `json:"old_password,omitempty"`
	Role           NullableUserRole `json:"role,omitempty"`
	Token          NullableString   `json:"token,omitempty"`
}

// NewUpdateUser instantiates a new UpdateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUser() *UpdateUser {
	this := UpdateUser{}
	return &this
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *UpdateUser) GetChangePassword() bool {
	if o == nil || IsNil(o.ChangePassword) {
		var ret bool
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetChangePasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ChangePassword) {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *UpdateUser) HasChangePassword() bool {
	if o != nil && !IsNil(o.ChangePassword) {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given bool and assigns it to the ChangePassword field.
func (o *UpdateUser) SetChangePassword(v bool) {
	o.ChangePassword = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateUser) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UpdateUser) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UpdateUser) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UpdateUser) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateUser) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UpdateUser) SetLastName(v string) {
	o.LastName.Set(&v)
}

// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UpdateUser) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UpdateUser) UnsetLastName() {
	o.LastName.Unset()
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword.Get()) {
		var ret string
		return ret
	}
	return *o.NewPassword.Get()
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewPassword.Get(), o.NewPassword.IsSet()
}

// HasNewPassword returns a boolean if a field has been set.
func (o *UpdateUser) HasNewPassword() bool {
	if o != nil && o.NewPassword.IsSet() {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given NullableString and assigns it to the NewPassword field.
func (o *UpdateUser) SetNewPassword(v string) {
	o.NewPassword.Set(&v)
}

// SetNewPasswordNil sets the value for NewPassword to be an explicit nil
func (o *UpdateUser) SetNewPasswordNil() {
	o.NewPassword.Set(nil)
}

// UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
func (o *UpdateUser) UnsetNewPassword() {
	o.NewPassword.Unset()
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetOldPassword() string {
	if o == nil || IsNil(o.OldPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OldPassword.Get()
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetOldPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldPassword.Get(), o.OldPassword.IsSet()
}

// HasOldPassword returns a boolean if a field has been set.
func (o *UpdateUser) HasOldPassword() bool {
	if o != nil && o.OldPassword.IsSet() {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given NullableString and assigns it to the OldPassword field.
func (o *UpdateUser) SetOldPassword(v string) {
	o.OldPassword.Set(&v)
}

// SetOldPasswordNil sets the value for OldPassword to be an explicit nil
func (o *UpdateUser) SetOldPasswordNil() {
	o.OldPassword.Set(nil)
}

// UnsetOldPassword ensures that no value is present for OldPassword, not even an explicit nil
func (o *UpdateUser) UnsetOldPassword() {
	o.OldPassword.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetRole() UserRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret UserRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetRoleOk() (*UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateUser) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableUserRole and assigns it to the Role field.
func (o *UpdateUser) SetRole(v UserRole) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *UpdateUser) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *UpdateUser) UnsetRole() {
	o.Role.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateUser) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *UpdateUser) SetToken(v string) {
	o.Token.Set(&v)
}

// SetTokenNil sets the value for Token to be an explicit nil
func (o *UpdateUser) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *UpdateUser) UnsetToken() {
	o.Token.Unset()
}

func (o UpdateUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangePassword) {
		toSerialize["change_password"] = o.ChangePassword
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.NewPassword.IsSet() {
		toSerialize["new_password"] = o.NewPassword.Get()
	}
	if o.OldPassword.IsSet() {
		toSerialize["old_password"] = o.OldPassword.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableUpdateUser struct {
	value *UpdateUser
	isSet bool
}

func (v NullableUpdateUser) Get() *UpdateUser {
	return v.value
}

func (v *NullableUpdateUser) Set(val *UpdateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUser(val *UpdateUser) *NullableUpdateUser {
	return &NullableUpdateUser{value: val, isSet: true}
}

func (v NullableUpdateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
