package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Layout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Layout{}

// Layout struct for Layout
type Layout struct {
	H       int64  `json:"h"`
	I       int64  `json:"i"`
	PanelId string `json:"panelId"`
	Static  bool   `json:"static"`
	W       int64  `json:"w"`
	X       int64  `json:"x"`
	Y       int64  `json:"y"`
}

type _Layout Layout

// NewLayout instantiates a new Layout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayout(h int64, i int64, panelId string, static bool, w int64, x int64, y int64) *Layout {
	this := Layout{}
	this.H = h
	this.I = i
	this.PanelId = panelId
	this.Static = static
	this.W = w
	this.X = x
	this.Y = y
	return &this
}

// GetH returns the H field value
func (o *Layout) GetH() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.H
}

// GetHOk returns a tuple with the H field value
// and a boolean to check if the value has been set.
func (o *Layout) GetHOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.H, true
}

// SetH sets field value
func (o *Layout) SetH(v int64) {
	o.H = v
}

// GetI returns the I field value
func (o *Layout) GetI() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.I
}

// GetIOk returns a tuple with the I field value
// and a boolean to check if the value has been set.
func (o *Layout) GetIOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.I, true
}

// SetI sets field value
func (o *Layout) SetI(v int64) {
	o.I = v
}

// GetPanelId returns the PanelId field value
func (o *Layout) GetPanelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PanelId
}

// GetPanelIdOk returns a tuple with the PanelId field value
// and a boolean to check if the value has been set.
func (o *Layout) GetPanelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PanelId, true
}

// SetPanelId sets field value
func (o *Layout) SetPanelId(v string) {
	o.PanelId = v
}

// GetStatic returns the Static field value
func (o *Layout) GetStatic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Static
}

// GetStaticOk returns a tuple with the Static field value
// and a boolean to check if the value has been set.
func (o *Layout) GetStaticOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Static, true
}

// SetStatic sets field value
func (o *Layout) SetStatic(v bool) {
	o.Static = v
}

// GetW returns the W field value
func (o *Layout) GetW() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *Layout) GetWOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *Layout) SetW(v int64) {
	o.W = v
}

// GetX returns the X field value
func (o *Layout) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *Layout) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *Layout) SetX(v int64) {
	o.X = v
}

// GetY returns the Y field value
func (o *Layout) GetY() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *Layout) GetYOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *Layout) SetY(v int64) {
	o.Y = v
}

func (o Layout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Layout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["h"] = o.H
	toSerialize["i"] = o.I
	toSerialize["panelId"] = o.PanelId
	toSerialize["static"] = o.Static
	toSerialize["w"] = o.W
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *Layout) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"h",
		"i",
		"panelId",
		"static",
		"w",
		"x",
		"y",
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

	varLayout := _Layout{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLayout)

	if err != nil {
		return err
	}

	*o = Layout(varLayout)

	return err
}

type NullableLayout struct {
	value *Layout
	isSet bool
}

func (v NullableLayout) Get() *Layout {
	return v.value
}

func (v *NullableLayout) Set(val *Layout) {
	v.value = val
	v.isSet = true
}

func (v NullableLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayout(val *Layout) *NullableLayout {
	return &NullableLayout{value: val, isSet: true}
}

func (v NullableLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
