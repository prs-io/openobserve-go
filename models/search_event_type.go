package models

import (
	"encoding/json"
	"fmt"
)

// SearchEventType the model 'SearchEventType'
type SearchEventType string

// List of SearchEventType
const (
	UI            SearchEventType = "ui"
	DASHBOARDS    SearchEventType = "dashboards"
	REPORTS       SearchEventType = "reports"
	ALERTS        SearchEventType = "alerts"
	VALUES        SearchEventType = "values"
	OTHER         SearchEventType = "other"
	RUM           SearchEventType = "rum"
	DERIVEDSTREAM SearchEventType = "derivedstream"
	SEARCHJOB     SearchEventType = "searchjob"
)

// All allowed values of SearchEventType enum
var AllowedSearchEventTypeEnumValues = []SearchEventType{
	"ui",
	"dashboards",
	"reports",
	"alerts",
	"values",
	"other",
	"rum",
	"derivedstream",
	"searchjob",
}

func (v *SearchEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchEventType(value)
	for _, existing := range AllowedSearchEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchEventType", value)
}

// NewSearchEventTypeFromValue returns a pointer to a valid SearchEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchEventTypeFromValue(v string) (*SearchEventType, error) {
	ev := SearchEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchEventType: valid values are %v", v, AllowedSearchEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchEventType) IsValid() bool {
	for _, existing := range AllowedSearchEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchEventType value
func (v SearchEventType) Ptr() *SearchEventType {
	return &v
}

type NullableSearchEventType struct {
	value *SearchEventType
	isSet bool
}

func (v NullableSearchEventType) Get() *SearchEventType {
	return v.value
}

func (v *NullableSearchEventType) Set(val *SearchEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEventType(val *SearchEventType) *NullableSearchEventType {
	return &NullableSearchEventType{value: val, isSet: true}
}

func (v NullableSearchEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
