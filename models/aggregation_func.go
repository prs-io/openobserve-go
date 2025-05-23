package models

import (
	"encoding/json"
	"fmt"
)

// AggregationFunc the model 'AggregationFunc'
type AggregationFunc string

// List of AggregationFunc
const (
	COUNT_DISTINCT AggregationFunc = "count-distinct"
)

// All allowed values of AggregationFunc enum
var AllowedAggregationFuncEnumValues = []AggregationFunc{
	"count",
	"count-distinct",
	"histogram",
	"sum",
	"min",
	"max",
	"avg",
	"median",
}

func (v *AggregationFunc) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AggregationFunc(value)
	for _, existing := range AllowedAggregationFuncEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AggregationFunc", value)
}

// NewAggregationFuncFromValue returns a pointer to a valid AggregationFunc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregationFuncFromValue(v string) (*AggregationFunc, error) {
	ev := AggregationFunc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AggregationFunc: valid values are %v", v, AllowedAggregationFuncEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregationFunc) IsValid() bool {
	for _, existing := range AllowedAggregationFuncEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationFunc value
func (v AggregationFunc) Ptr() *AggregationFunc {
	return &v
}

type NullableAggregationFunc struct {
	value *AggregationFunc
	isSet bool
}

func (v NullableAggregationFunc) Get() *AggregationFunc {
	return v.value
}

func (v *NullableAggregationFunc) Set(val *AggregationFunc) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationFunc) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationFunc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationFunc(val *AggregationFunc) *NullableAggregationFunc {
	return &NullableAggregationFunc{value: val, isSet: true}
}

func (v NullableAggregationFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationFunc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
