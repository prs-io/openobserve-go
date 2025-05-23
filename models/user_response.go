package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponse{}

// UserResponse struct for UserResponse
type UserResponse struct {
	Email      string   `json:"email"`
	FirstName  *string  `json:"first_name,omitempty"`
	IsExternal *bool    `json:"is_external,omitempty"`
	LastName   *string  `json:"last_name,omitempty"`
	Role       UserRole `json:"role"`
}

type _UserResponse UserResponse

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse(email string, role UserRole) *UserResponse {
	this := UserResponse{}
	this.Email = email
	this.Role = role
	return &this
}

// GetEmail returns the Email field value
func (o *UserResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserResponse) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *UserResponse) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIsExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternal) {
		return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *UserResponse) HasIsExternal() bool {
	if o != nil && !IsNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *UserResponse) SetIsExternal(v bool) {
	o.IsExternal = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetRole returns the Role field value
func (o *UserResponse) GetRole() UserRole {
	if o == nil {
		var ret UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetRoleOk() (*UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserResponse) SetRole(v UserRole) {
	o.Role = v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.IsExternal) {
		toSerialize["is_external"] = o.IsExternal
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *UserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUserResponse := _UserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserResponse)

	if err != nil {
		return err
	}

	*o = UserResponse(varUserResponse)

	return err
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
