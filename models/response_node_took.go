package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ResponseNodeTook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseNodeTook{}

// ResponseNodeTook struct for ResponseNodeTook
type ResponseNodeTook struct {
	IsIngester bool   `json:"is_ingester"`
	Node       string `json:"node"`
	Took       int32  `json:"took"`
}

type _ResponseNodeTook ResponseNodeTook

// NewResponseNodeTook instantiates a new ResponseNodeTook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseNodeTook(isIngester bool, node string, took int32) *ResponseNodeTook {
	this := ResponseNodeTook{}
	this.IsIngester = isIngester
	this.Node = node
	this.Took = took
	return &this
}

// GetIsIngester returns the IsIngester field value
func (o *ResponseNodeTook) GetIsIngester() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIngester
}

// GetIsIngesterOk returns a tuple with the IsIngester field value
// and a boolean to check if the value has been set.
func (o *ResponseNodeTook) GetIsIngesterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsIngester, true
}

// SetIsIngester sets field value
func (o *ResponseNodeTook) SetIsIngester(v bool) {
	o.IsIngester = v
}

// GetNode returns the Node field value
func (o *ResponseNodeTook) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *ResponseNodeTook) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *ResponseNodeTook) SetNode(v string) {
	o.Node = v
}

// GetTook returns the Took field value
func (o *ResponseNodeTook) GetTook() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Took
}

// GetTookOk returns a tuple with the Took field value
// and a boolean to check if the value has been set.
func (o *ResponseNodeTook) GetTookOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Took, true
}

// SetTook sets field value
func (o *ResponseNodeTook) SetTook(v int32) {
	o.Took = v
}

func (o ResponseNodeTook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseNodeTook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_ingester"] = o.IsIngester
	toSerialize["node"] = o.Node
	toSerialize["took"] = o.Took
	return toSerialize, nil
}

func (o *ResponseNodeTook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_ingester",
		"node",
		"took",
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

	varResponseNodeTook := _ResponseNodeTook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseNodeTook)

	if err != nil {
		return err
	}

	*o = ResponseNodeTook(varResponseNodeTook)

	return err
}

type NullableResponseNodeTook struct {
	value *ResponseNodeTook
	isSet bool
}

func (v NullableResponseNodeTook) Get() *ResponseNodeTook {
	return v.value
}

func (v *NullableResponseNodeTook) Set(val *ResponseNodeTook) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseNodeTook) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseNodeTook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseNodeTook(val *ResponseNodeTook) *NullableResponseNodeTook {
	return &NullableResponseNodeTook{value: val, isSet: true}
}

func (v NullableResponseNodeTook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseNodeTook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
