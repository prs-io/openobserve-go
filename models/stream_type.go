package models

import (
	"encoding/json"
	"fmt"
)

// StreamType the model 'StreamType'
type StreamType string

// List of StreamType
const (
	LOGS              StreamType = "logs"
	METRICS           StreamType = "metrics"
	TRACES            StreamType = "traces"
	ENRICHMENT_TABLES StreamType = "enrichment_tables"
	FILE_LIST         StreamType = "file_list"
	METADATA          StreamType = "metadata"
	INDEX             StreamType = "index"
)

// All allowed values of StreamType enum
var AllowedStreamTypeEnumValues = []StreamType{
	"logs",
	"metrics",
	"traces",
	"enrichment_tables",
	"file_list",
	"metadata",
	"index",
}

func (v *StreamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StreamType(value)
	for _, existing := range AllowedStreamTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StreamType", value)
}

// NewStreamTypeFromValue returns a pointer to a valid StreamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStreamTypeFromValue(v string) (*StreamType, error) {
	ev := StreamType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StreamType: valid values are %v", v, AllowedStreamTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StreamType) IsValid() bool {
	for _, existing := range AllowedStreamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StreamType value
func (v StreamType) Ptr() *StreamType {
	return &v
}

type NullableStreamType struct {
	value *StreamType
	isSet bool
}

func (v NullableStreamType) Get() *StreamType {
	return v.value
}

func (v *NullableStreamType) Set(val *StreamType) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamType) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamType(val *StreamType) *NullableStreamType {
	return &NullableStreamType{value: val, isSet: true}
}

func (v NullableStreamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
