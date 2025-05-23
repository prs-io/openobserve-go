package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AxisItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxisItem{}

// AxisItem struct for AxisItem
type AxisItem struct {
	AggregationFunction NullableAggregationFunc `json:"aggregationFunction,omitempty"`
	Alias               string                  `json:"alias"`
	Color               NullableString          `json:"color,omitempty"`
	Column              string                  `json:"column"`
	Label               string                  `json:"label"`
}

type _AxisItem AxisItem

// NewAxisItem instantiates a new AxisItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxisItem(alias string, column string, label string) *AxisItem {
	this := AxisItem{}
	this.Alias = alias
	this.Column = column
	this.Label = label
	return &this
}

// GetAggregationFunction returns the AggregationFunction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AxisItem) GetAggregationFunction() AggregationFunc {
	if o == nil || IsNil(o.AggregationFunction.Get()) {
		var ret AggregationFunc
		return ret
	}
	return *o.AggregationFunction.Get()
}

// GetAggregationFunctionOk returns a tuple with the AggregationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxisItem) GetAggregationFunctionOk() (*AggregationFunc, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationFunction.Get(), o.AggregationFunction.IsSet()
}

// HasAggregationFunction returns a boolean if a field has been set.
func (o *AxisItem) HasAggregationFunction() bool {
	if o != nil && o.AggregationFunction.IsSet() {
		return true
	}

	return false
}

// SetAggregationFunction gets a reference to the given NullableAggregationFunc and assigns it to the AggregationFunction field.
func (o *AxisItem) SetAggregationFunction(v AggregationFunc) {
	o.AggregationFunction.Set(&v)
}

// SetAggregationFunctionNil sets the value for AggregationFunction to be an explicit nil
func (o *AxisItem) SetAggregationFunctionNil() {
	o.AggregationFunction.Set(nil)
}

// UnsetAggregationFunction ensures that no value is present for AggregationFunction, not even an explicit nil
func (o *AxisItem) UnsetAggregationFunction() {
	o.AggregationFunction.Unset()
}

// GetAlias returns the Alias field value
func (o *AxisItem) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *AxisItem) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *AxisItem) SetAlias(v string) {
	o.Alias = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AxisItem) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxisItem) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *AxisItem) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *AxisItem) SetColor(v string) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *AxisItem) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *AxisItem) UnsetColor() {
	o.Color.Unset()
}

// GetColumn returns the Column field value
func (o *AxisItem) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *AxisItem) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *AxisItem) SetColumn(v string) {
	o.Column = v
}

// GetLabel returns the Label field value
func (o *AxisItem) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AxisItem) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AxisItem) SetLabel(v string) {
	o.Label = v
}

func (o AxisItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxisItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregationFunction.IsSet() {
		toSerialize["aggregationFunction"] = o.AggregationFunction.Get()
	}
	toSerialize["alias"] = o.Alias
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	toSerialize["column"] = o.Column
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *AxisItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"column",
		"label",
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

	varAxisItem := _AxisItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAxisItem)

	if err != nil {
		return err
	}

	*o = AxisItem(varAxisItem)

	return err
}

type NullableAxisItem struct {
	value *AxisItem
	isSet bool
}

func (v NullableAxisItem) Get() *AxisItem {
	return v.value
}

func (v *NullableAxisItem) Set(val *AxisItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAxisItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAxisItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxisItem(val *AxisItem) *NullableAxisItem {
	return &NullableAxisItem{value: val, isSet: true}
}

func (v NullableAxisItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxisItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
