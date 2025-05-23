package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HttpResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpResponse{}

// HttpResponse HTTP response code 200 is success code 400 is error code 404 is not found code 500 is internal server error code 503 is service unavailable code >= 1000 is custom error code message is the message or error message
type HttpResponse struct {
	Code        int32          `json:"code"`
	ErrorDetail NullableString `json:"error_detail,omitempty"`
	Message     string         `json:"message"`
	TraceId     NullableString `json:"trace_id,omitempty"`
}

type _HttpResponse HttpResponse

// NewHttpResponse instantiates a new HttpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpResponse(code int32, message string) *HttpResponse {
	this := HttpResponse{}
	this.Code = code
	this.Message = message
	return &this
}

// GetCode returns the Code field value
func (o *HttpResponse) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *HttpResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *HttpResponse) SetCode(v int32) {
	o.Code = v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpResponse) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpResponse) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *HttpResponse) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *HttpResponse) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}

// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *HttpResponse) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *HttpResponse) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetMessage returns the Message field value
func (o *HttpResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *HttpResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *HttpResponse) SetMessage(v string) {
	o.Message = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpResponse) GetTraceId() string {
	if o == nil || IsNil(o.TraceId.Get()) {
		var ret string
		return ret
	}
	return *o.TraceId.Get()
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpResponse) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceId.Get(), o.TraceId.IsSet()
}

// HasTraceId returns a boolean if a field has been set.
func (o *HttpResponse) HasTraceId() bool {
	if o != nil && o.TraceId.IsSet() {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given NullableString and assigns it to the TraceId field.
func (o *HttpResponse) SetTraceId(v string) {
	o.TraceId.Set(&v)
}

// SetTraceIdNil sets the value for TraceId to be an explicit nil
func (o *HttpResponse) SetTraceIdNil() {
	o.TraceId.Set(nil)
}

// UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil
func (o *HttpResponse) UnsetTraceId() {
	o.TraceId.Unset()
}

func (o HttpResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if o.ErrorDetail.IsSet() {
		toSerialize["error_detail"] = o.ErrorDetail.Get()
	}
	toSerialize["message"] = o.Message
	if o.TraceId.IsSet() {
		toSerialize["trace_id"] = o.TraceId.Get()
	}
	return toSerialize, nil
}

func (o *HttpResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varHttpResponse := _HttpResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHttpResponse)

	if err != nil {
		return err
	}

	*o = HttpResponse(varHttpResponse)

	return err
}

type NullableHttpResponse struct {
	value *HttpResponse
	isSet bool
}

func (v NullableHttpResponse) Get() *HttpResponse {
	return v.value
}

func (v *NullableHttpResponse) Set(val *HttpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpResponse(val *HttpResponse) *NullableHttpResponse {
	return &NullableHttpResponse{value: val, isSet: true}
}

func (v NullableHttpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
