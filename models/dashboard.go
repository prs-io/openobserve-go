package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Dashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboard{}

// Dashboard struct for Dashboard
type Dashboard struct {
	Created     *time.Time        `json:"created,omitempty"`
	DashboardId *string           `json:"dashboardId,omitempty"`
	Description string            `json:"description"`
	Layouts     []Layout          `json:"layouts,omitempty"`
	Owner       *string           `json:"owner,omitempty"`
	Panels      []Panel           `json:"panels,omitempty"`
	Role        *string           `json:"role,omitempty"`
	Title       string            `json:"title"`
	Variables   NullableVariables `json:"variables,omitempty"`
}

type _Dashboard Dashboard

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard(description string, title string) *Dashboard {
	this := Dashboard{}
	this.Description = description
	this.Title = title
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Dashboard) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Dashboard) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Dashboard) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise.
func (o *Dashboard) GetDashboardId() string {
	if o == nil || IsNil(o.DashboardId) {
		var ret string
		return ret
	}
	return *o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDashboardIdOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardId) {
		return nil, false
	}
	return o.DashboardId, true
}

// HasDashboardId returns a boolean if a field has been set.
func (o *Dashboard) HasDashboardId() bool {
	if o != nil && !IsNil(o.DashboardId) {
		return true
	}

	return false
}

// SetDashboardId gets a reference to the given string and assigns it to the DashboardId field.
func (o *Dashboard) SetDashboardId(v string) {
	o.DashboardId = &v
}

// GetDescription returns the Description field value
func (o *Dashboard) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Dashboard) SetDescription(v string) {
	o.Description = v
}

// GetLayouts returns the Layouts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetLayouts() []Layout {
	if o == nil {
		var ret []Layout
		return ret
	}
	return o.Layouts
}

// GetLayoutsOk returns a tuple with the Layouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetLayoutsOk() ([]Layout, bool) {
	if o == nil || IsNil(o.Layouts) {
		return nil, false
	}
	return o.Layouts, true
}

// HasLayouts returns a boolean if a field has been set.
func (o *Dashboard) HasLayouts() bool {
	if o != nil && !IsNil(o.Layouts) {
		return true
	}

	return false
}

// SetLayouts gets a reference to the given []Layout and assigns it to the Layouts field.
func (o *Dashboard) SetLayouts(v []Layout) {
	o.Layouts = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Dashboard) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Dashboard) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Dashboard) SetOwner(v string) {
	o.Owner = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *Dashboard) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *Dashboard) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *Dashboard) SetPanels(v []Panel) {
	o.Panels = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Dashboard) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Dashboard) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Dashboard) SetRole(v string) {
	o.Role = &v
}

// GetTitle returns the Title field value
func (o *Dashboard) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Dashboard) SetTitle(v string) {
	o.Title = v
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetVariables() Variables {
	if o == nil || IsNil(o.Variables.Get()) {
		var ret Variables
		return ret
	}
	return *o.Variables.Get()
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetVariablesOk() (*Variables, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variables.Get(), o.Variables.IsSet()
}

// HasVariables returns a boolean if a field has been set.
func (o *Dashboard) HasVariables() bool {
	if o != nil && o.Variables.IsSet() {
		return true
	}

	return false
}

// SetVariables gets a reference to the given NullableVariables and assigns it to the Variables field.
func (o *Dashboard) SetVariables(v Variables) {
	o.Variables.Set(&v)
}

// SetVariablesNil sets the value for Variables to be an explicit nil
func (o *Dashboard) SetVariablesNil() {
	o.Variables.Set(nil)
}

// UnsetVariables ensures that no value is present for Variables, not even an explicit nil
func (o *Dashboard) UnsetVariables() {
	o.Variables.Unset()
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DashboardId) {
		toSerialize["dashboardId"] = o.DashboardId
	}
	toSerialize["description"] = o.Description
	if o.Layouts != nil {
		toSerialize["layouts"] = o.Layouts
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Panels) {
		toSerialize["panels"] = o.Panels
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	toSerialize["title"] = o.Title
	if o.Variables.IsSet() {
		toSerialize["variables"] = o.Variables.Get()
	}
	return toSerialize, nil
}

func (o *Dashboard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"title",
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

	varDashboard := _Dashboard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboard)

	if err != nil {
		return err
	}

	*o = Dashboard(varDashboard)

	return err
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
