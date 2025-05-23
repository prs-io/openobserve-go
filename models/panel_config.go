package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PanelConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PanelConfig{}

// PanelConfig struct for PanelConfig
type PanelConfig struct {
	Description     string         `json:"description"`
	LegendsPosition NullableString `json:"legends_position,omitempty"`
	PromqlLegend    NullableString `json:"promql_legend,omitempty"`
	ShowLegends     bool           `json:"show_legends"`
	Title           string         `json:"title"`
	Unit            NullableString `json:"unit,omitempty"`
	UnitCustom      NullableString `json:"unit_custom,omitempty"`
}

type _PanelConfig PanelConfig

// NewPanelConfig instantiates a new PanelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanelConfig(description string, showLegends bool, title string) *PanelConfig {
	this := PanelConfig{}
	this.Description = description
	this.ShowLegends = showLegends
	this.Title = title
	return &this
}

// GetDescription returns the Description field value
func (o *PanelConfig) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PanelConfig) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PanelConfig) SetDescription(v string) {
	o.Description = v
}

// GetLegendsPosition returns the LegendsPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelConfig) GetLegendsPosition() string {
	if o == nil || IsNil(o.LegendsPosition.Get()) {
		var ret string
		return ret
	}
	return *o.LegendsPosition.Get()
}

// GetLegendsPositionOk returns a tuple with the LegendsPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelConfig) GetLegendsPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegendsPosition.Get(), o.LegendsPosition.IsSet()
}

// HasLegendsPosition returns a boolean if a field has been set.
func (o *PanelConfig) HasLegendsPosition() bool {
	if o != nil && o.LegendsPosition.IsSet() {
		return true
	}

	return false
}

// SetLegendsPosition gets a reference to the given NullableString and assigns it to the LegendsPosition field.
func (o *PanelConfig) SetLegendsPosition(v string) {
	o.LegendsPosition.Set(&v)
}

// SetLegendsPositionNil sets the value for LegendsPosition to be an explicit nil
func (o *PanelConfig) SetLegendsPositionNil() {
	o.LegendsPosition.Set(nil)
}

// UnsetLegendsPosition ensures that no value is present for LegendsPosition, not even an explicit nil
func (o *PanelConfig) UnsetLegendsPosition() {
	o.LegendsPosition.Unset()
}

// GetPromqlLegend returns the PromqlLegend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelConfig) GetPromqlLegend() string {
	if o == nil || IsNil(o.PromqlLegend.Get()) {
		var ret string
		return ret
	}
	return *o.PromqlLegend.Get()
}

// GetPromqlLegendOk returns a tuple with the PromqlLegend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelConfig) GetPromqlLegendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromqlLegend.Get(), o.PromqlLegend.IsSet()
}

// HasPromqlLegend returns a boolean if a field has been set.
func (o *PanelConfig) HasPromqlLegend() bool {
	if o != nil && o.PromqlLegend.IsSet() {
		return true
	}

	return false
}

// SetPromqlLegend gets a reference to the given NullableString and assigns it to the PromqlLegend field.
func (o *PanelConfig) SetPromqlLegend(v string) {
	o.PromqlLegend.Set(&v)
}

// SetPromqlLegendNil sets the value for PromqlLegend to be an explicit nil
func (o *PanelConfig) SetPromqlLegendNil() {
	o.PromqlLegend.Set(nil)
}

// UnsetPromqlLegend ensures that no value is present for PromqlLegend, not even an explicit nil
func (o *PanelConfig) UnsetPromqlLegend() {
	o.PromqlLegend.Unset()
}

// GetShowLegends returns the ShowLegends field value
func (o *PanelConfig) GetShowLegends() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowLegends
}

// GetShowLegendsOk returns a tuple with the ShowLegends field value
// and a boolean to check if the value has been set.
func (o *PanelConfig) GetShowLegendsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowLegends, true
}

// SetShowLegends sets field value
func (o *PanelConfig) SetShowLegends(v bool) {
	o.ShowLegends = v
}

// GetTitle returns the Title field value
func (o *PanelConfig) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PanelConfig) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PanelConfig) SetTitle(v string) {
	o.Title = v
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelConfig) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelConfig) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *PanelConfig) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *PanelConfig) SetUnit(v string) {
	o.Unit.Set(&v)
}

// SetUnitNil sets the value for Unit to be an explicit nil
func (o *PanelConfig) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *PanelConfig) UnsetUnit() {
	o.Unit.Unset()
}

// GetUnitCustom returns the UnitCustom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanelConfig) GetUnitCustom() string {
	if o == nil || IsNil(o.UnitCustom.Get()) {
		var ret string
		return ret
	}
	return *o.UnitCustom.Get()
}

// GetUnitCustomOk returns a tuple with the UnitCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanelConfig) GetUnitCustomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitCustom.Get(), o.UnitCustom.IsSet()
}

// HasUnitCustom returns a boolean if a field has been set.
func (o *PanelConfig) HasUnitCustom() bool {
	if o != nil && o.UnitCustom.IsSet() {
		return true
	}

	return false
}

// SetUnitCustom gets a reference to the given NullableString and assigns it to the UnitCustom field.
func (o *PanelConfig) SetUnitCustom(v string) {
	o.UnitCustom.Set(&v)
}

// SetUnitCustomNil sets the value for UnitCustom to be an explicit nil
func (o *PanelConfig) SetUnitCustomNil() {
	o.UnitCustom.Set(nil)
}

// UnsetUnitCustom ensures that no value is present for UnitCustom, not even an explicit nil
func (o *PanelConfig) UnsetUnitCustom() {
	o.UnitCustom.Unset()
}

func (o PanelConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PanelConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if o.LegendsPosition.IsSet() {
		toSerialize["legends_position"] = o.LegendsPosition.Get()
	}
	if o.PromqlLegend.IsSet() {
		toSerialize["promql_legend"] = o.PromqlLegend.Get()
	}
	toSerialize["show_legends"] = o.ShowLegends
	toSerialize["title"] = o.Title
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.UnitCustom.IsSet() {
		toSerialize["unit_custom"] = o.UnitCustom.Get()
	}
	return toSerialize, nil
}

func (o *PanelConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"show_legends",
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

	varPanelConfig := _PanelConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPanelConfig)

	if err != nil {
		return err
	}

	*o = PanelConfig(varPanelConfig)

	return err
}

type NullablePanelConfig struct {
	value *PanelConfig
	isSet bool
}

func (v NullablePanelConfig) Get() *PanelConfig {
	return v.value
}

func (v *NullablePanelConfig) Set(val *PanelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePanelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePanelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanelConfig(val *PanelConfig) *NullablePanelConfig {
	return &NullablePanelConfig{value: val, isSet: true}
}

func (v NullablePanelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
