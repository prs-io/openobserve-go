package models

import (
	"encoding/json"
)

// checks if the SearchEventContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEventContext{}

// SearchEventContext struct for SearchEventContext
type SearchEventContext struct {
	AlertKey         NullableString `json:"alert_key,omitempty"`
	DashboardId      NullableString `json:"dashboard_id,omitempty"`
	DashboardName    NullableString `json:"dashboard_name,omitempty"`
	DerivedStreamKey NullableString `json:"derived_stream_key,omitempty"`
	FolderId         NullableString `json:"folder_id,omitempty"`
	FolderName       NullableString `json:"folder_name,omitempty"`
	ReportId         NullableString `json:"report_id,omitempty"`
}

// NewSearchEventContext instantiates a new SearchEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEventContext() *SearchEventContext {
	this := SearchEventContext{}
	return &this
}

// GetAlertKey returns the AlertKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetAlertKey() string {
	if o == nil || IsNil(o.AlertKey.Get()) {
		var ret string
		return ret
	}
	return *o.AlertKey.Get()
}

// GetAlertKeyOk returns a tuple with the AlertKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetAlertKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertKey.Get(), o.AlertKey.IsSet()
}

// HasAlertKey returns a boolean if a field has been set.
func (o *SearchEventContext) HasAlertKey() bool {
	if o != nil && o.AlertKey.IsSet() {
		return true
	}

	return false
}

// SetAlertKey gets a reference to the given NullableString and assigns it to the AlertKey field.
func (o *SearchEventContext) SetAlertKey(v string) {
	o.AlertKey.Set(&v)
}

// SetAlertKeyNil sets the value for AlertKey to be an explicit nil
func (o *SearchEventContext) SetAlertKeyNil() {
	o.AlertKey.Set(nil)
}

// UnsetAlertKey ensures that no value is present for AlertKey, not even an explicit nil
func (o *SearchEventContext) UnsetAlertKey() {
	o.AlertKey.Unset()
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetDashboardId() string {
	if o == nil || IsNil(o.DashboardId.Get()) {
		var ret string
		return ret
	}
	return *o.DashboardId.Get()
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardId.Get(), o.DashboardId.IsSet()
}

// HasDashboardId returns a boolean if a field has been set.
func (o *SearchEventContext) HasDashboardId() bool {
	if o != nil && o.DashboardId.IsSet() {
		return true
	}

	return false
}

// SetDashboardId gets a reference to the given NullableString and assigns it to the DashboardId field.
func (o *SearchEventContext) SetDashboardId(v string) {
	o.DashboardId.Set(&v)
}

// SetDashboardIdNil sets the value for DashboardId to be an explicit nil
func (o *SearchEventContext) SetDashboardIdNil() {
	o.DashboardId.Set(nil)
}

// UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
func (o *SearchEventContext) UnsetDashboardId() {
	o.DashboardId.Unset()
}

// GetDashboardName returns the DashboardName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetDashboardName() string {
	if o == nil || IsNil(o.DashboardName.Get()) {
		var ret string
		return ret
	}
	return *o.DashboardName.Get()
}

// GetDashboardNameOk returns a tuple with the DashboardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetDashboardNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardName.Get(), o.DashboardName.IsSet()
}

// HasDashboardName returns a boolean if a field has been set.
func (o *SearchEventContext) HasDashboardName() bool {
	if o != nil && o.DashboardName.IsSet() {
		return true
	}

	return false
}

// SetDashboardName gets a reference to the given NullableString and assigns it to the DashboardName field.
func (o *SearchEventContext) SetDashboardName(v string) {
	o.DashboardName.Set(&v)
}

// SetDashboardNameNil sets the value for DashboardName to be an explicit nil
func (o *SearchEventContext) SetDashboardNameNil() {
	o.DashboardName.Set(nil)
}

// UnsetDashboardName ensures that no value is present for DashboardName, not even an explicit nil
func (o *SearchEventContext) UnsetDashboardName() {
	o.DashboardName.Unset()
}

// GetDerivedStreamKey returns the DerivedStreamKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetDerivedStreamKey() string {
	if o == nil || IsNil(o.DerivedStreamKey.Get()) {
		var ret string
		return ret
	}
	return *o.DerivedStreamKey.Get()
}

// GetDerivedStreamKeyOk returns a tuple with the DerivedStreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetDerivedStreamKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DerivedStreamKey.Get(), o.DerivedStreamKey.IsSet()
}

// HasDerivedStreamKey returns a boolean if a field has been set.
func (o *SearchEventContext) HasDerivedStreamKey() bool {
	if o != nil && o.DerivedStreamKey.IsSet() {
		return true
	}

	return false
}

// SetDerivedStreamKey gets a reference to the given NullableString and assigns it to the DerivedStreamKey field.
func (o *SearchEventContext) SetDerivedStreamKey(v string) {
	o.DerivedStreamKey.Set(&v)
}

// SetDerivedStreamKeyNil sets the value for DerivedStreamKey to be an explicit nil
func (o *SearchEventContext) SetDerivedStreamKeyNil() {
	o.DerivedStreamKey.Set(nil)
}

// UnsetDerivedStreamKey ensures that no value is present for DerivedStreamKey, not even an explicit nil
func (o *SearchEventContext) UnsetDerivedStreamKey() {
	o.DerivedStreamKey.Unset()
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *SearchEventContext) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *SearchEventContext) SetFolderId(v string) {
	o.FolderId.Set(&v)
}

// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *SearchEventContext) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *SearchEventContext) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetFolderName returns the FolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetFolderName() string {
	if o == nil || IsNil(o.FolderName.Get()) {
		var ret string
		return ret
	}
	return *o.FolderName.Get()
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetFolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderName.Get(), o.FolderName.IsSet()
}

// HasFolderName returns a boolean if a field has been set.
func (o *SearchEventContext) HasFolderName() bool {
	if o != nil && o.FolderName.IsSet() {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given NullableString and assigns it to the FolderName field.
func (o *SearchEventContext) SetFolderName(v string) {
	o.FolderName.Set(&v)
}

// SetFolderNameNil sets the value for FolderName to be an explicit nil
func (o *SearchEventContext) SetFolderNameNil() {
	o.FolderName.Set(nil)
}

// UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
func (o *SearchEventContext) UnsetFolderName() {
	o.FolderName.Unset()
}

// GetReportId returns the ReportId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchEventContext) GetReportId() string {
	if o == nil || IsNil(o.ReportId.Get()) {
		var ret string
		return ret
	}
	return *o.ReportId.Get()
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchEventContext) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportId.Get(), o.ReportId.IsSet()
}

// HasReportId returns a boolean if a field has been set.
func (o *SearchEventContext) HasReportId() bool {
	if o != nil && o.ReportId.IsSet() {
		return true
	}

	return false
}

// SetReportId gets a reference to the given NullableString and assigns it to the ReportId field.
func (o *SearchEventContext) SetReportId(v string) {
	o.ReportId.Set(&v)
}

// SetReportIdNil sets the value for ReportId to be an explicit nil
func (o *SearchEventContext) SetReportIdNil() {
	o.ReportId.Set(nil)
}

// UnsetReportId ensures that no value is present for ReportId, not even an explicit nil
func (o *SearchEventContext) UnsetReportId() {
	o.ReportId.Unset()
}

func (o SearchEventContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEventContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertKey.IsSet() {
		toSerialize["alert_key"] = o.AlertKey.Get()
	}
	if o.DashboardId.IsSet() {
		toSerialize["dashboard_id"] = o.DashboardId.Get()
	}
	if o.DashboardName.IsSet() {
		toSerialize["dashboard_name"] = o.DashboardName.Get()
	}
	if o.DerivedStreamKey.IsSet() {
		toSerialize["derived_stream_key"] = o.DerivedStreamKey.Get()
	}
	if o.FolderId.IsSet() {
		toSerialize["folder_id"] = o.FolderId.Get()
	}
	if o.FolderName.IsSet() {
		toSerialize["folder_name"] = o.FolderName.Get()
	}
	if o.ReportId.IsSet() {
		toSerialize["report_id"] = o.ReportId.Get()
	}
	return toSerialize, nil
}

type NullableSearchEventContext struct {
	value *SearchEventContext
	isSet bool
}

func (v NullableSearchEventContext) Get() *SearchEventContext {
	return v.value
}

func (v *NullableSearchEventContext) Set(val *SearchEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEventContext(val *SearchEventContext) *NullableSearchEventContext {
	return &NullableSearchEventContext{value: val, isSet: true}
}

func (v NullableSearchEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
