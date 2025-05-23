package models

import (
	"encoding/json"
	"fmt"
)

// OrderBy the model 'OrderBy'
type OrderBy string

// List of OrderBy
const (
	DESC OrderBy = "desc"
	ASC  OrderBy = "asc"
)

// All allowed values of OrderBy enum
var AllowedOrderByEnumValues = []OrderBy{
	"desc",
	"asc",
}

func (v *OrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderBy(value)
	for _, existing := range AllowedOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderBy", value)
}

// NewOrderByFromValue returns a pointer to a valid OrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderByFromValue(v string) (*OrderBy, error) {
	ev := OrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderBy: valid values are %v", v, AllowedOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderBy) IsValid() bool {
	for _, existing := range AllowedOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderBy value
func (v OrderBy) Ptr() *OrderBy {
	return &v
}

type NullableOrderBy struct {
	value *OrderBy
	isSet bool
}

func (v NullableOrderBy) Get() *OrderBy {
	return v.value
}

func (v *NullableOrderBy) Set(val *OrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBy(val *OrderBy) *NullableOrderBy {
	return &NullableOrderBy{value: val, isSet: true}
}

func (v NullableOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
