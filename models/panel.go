package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Panel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Panel{}

// Panel struct for Panel
type Panel struct {
	Config      PanelConfig `json:"config"`
	CustomQuery bool        `json:"customQuery"`
	Fields      PanelFields `json:"fields"`
	Id          string      `json:"id"`
	Query       string      `json:"query"`
	QueryType   *string     `json:"queryType,omitempty"`
	Type        string      `json:"type"`
}

type _Panel Panel

// NewPanel instantiates a new Panel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanel(config PanelConfig, customQuery bool, fields PanelFields, id string, query string, type_ string) *Panel {
	this := Panel{}
	this.Config = config
	this.CustomQuery = customQuery
	this.Fields = fields
	this.Id = id
	this.Query = query
	this.Type = type_
	return &this
}

// GetConfig returns the Config field value
func (o *Panel) GetConfig() PanelConfig {
	if o == nil {
		var ret PanelConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Panel) GetConfigOk() (*PanelConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Panel) SetConfig(v PanelConfig) {
	o.Config = v
}

// GetCustomQuery returns the CustomQuery field value
func (o *Panel) GetCustomQuery() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CustomQuery
}

// GetCustomQueryOk returns a tuple with the CustomQuery field value
// and a boolean to check if the value has been set.
func (o *Panel) GetCustomQueryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomQuery, true
}

// SetCustomQuery sets field value
func (o *Panel) SetCustomQuery(v bool) {
	o.CustomQuery = v
}

// GetFields returns the Fields field value
func (o *Panel) GetFields() PanelFields {
	if o == nil {
		var ret PanelFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *Panel) GetFieldsOk() (*PanelFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *Panel) SetFields(v PanelFields) {
	o.Fields = v
}

// GetId returns the Id field value
func (o *Panel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Panel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Panel) SetId(v string) {
	o.Id = v
}

// GetQuery returns the Query field value
func (o *Panel) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Panel) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Panel) SetQuery(v string) {
	o.Query = v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *Panel) GetQueryType() string {
	if o == nil || IsNil(o.QueryType) {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Panel) GetQueryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryType) {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *Panel) HasQueryType() bool {
	if o != nil && !IsNil(o.QueryType) {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *Panel) SetQueryType(v string) {
	o.QueryType = &v
}

// GetType returns the Type field value
func (o *Panel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Panel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Panel) SetType(v string) {
	o.Type = v
}

func (o Panel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Panel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["customQuery"] = o.CustomQuery
	toSerialize["fields"] = o.Fields
	toSerialize["id"] = o.Id
	toSerialize["query"] = o.Query
	if !IsNil(o.QueryType) {
		toSerialize["queryType"] = o.QueryType
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *Panel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
		"customQuery",
		"fields",
		"id",
		"query",
		"type",
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

	varPanel := _Panel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPanel)

	if err != nil {
		return err
	}

	*o = Panel(varPanel)

	return err
}

type NullablePanel struct {
	value *Panel
	isSet bool
}

func (v NullablePanel) Get() *Panel {
	return v.value
}

func (v *NullablePanel) Set(val *Panel) {
	v.value = val
	v.isSet = true
}

func (v NullablePanel) IsSet() bool {
	return v.isSet
}

func (v *NullablePanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanel(val *Panel) *NullablePanel {
	return &NullablePanel{value: val, isSet: true}
}

func (v NullablePanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
