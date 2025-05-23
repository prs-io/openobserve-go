package models

import (
	"encoding/json"
	"fmt"
)

// RequestEncoding the model 'RequestEncoding'
type RequestEncoding string

// List of RequestEncoding
const (
	BASE64 RequestEncoding = "base64"
	EMPTY  RequestEncoding = "Empty"
)

// All allowed values of RequestEncoding enum
var AllowedRequestEncodingEnumValues = []RequestEncoding{
	"base64",
	"Empty",
}

func (v *RequestEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestEncoding(value)
	for _, existing := range AllowedRequestEncodingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestEncoding", value)
}

// NewRequestEncodingFromValue returns a pointer to a valid RequestEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestEncodingFromValue(v string) (*RequestEncoding, error) {
	ev := RequestEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestEncoding: valid values are %v", v, AllowedRequestEncodingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestEncoding) IsValid() bool {
	for _, existing := range AllowedRequestEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestEncoding value
func (v RequestEncoding) Ptr() *RequestEncoding {
	return &v
}

type NullableRequestEncoding struct {
	value *RequestEncoding
	isSet bool
}

func (v NullableRequestEncoding) Get() *RequestEncoding {
	return v.value
}

func (v *NullableRequestEncoding) Set(val *RequestEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestEncoding(val *RequestEncoding) *NullableRequestEncoding {
	return &NullableRequestEncoding{value: val, isSet: true}
}

func (v NullableRequestEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
