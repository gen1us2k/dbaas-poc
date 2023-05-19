/*
Percona Everest schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequest{}

// BadRequest Bad request response
type BadRequest struct {
	Message *string `json:"message,omitempty"`
}

// NewBadRequest instantiates a new BadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequest() *BadRequest {
	this := BadRequest{}
	return &this
}

// NewBadRequestWithDefaults instantiates a new BadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestWithDefaults() *BadRequest {
	this := BadRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BadRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BadRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BadRequest) SetMessage(v string) {
	o.Message = &v
}

func (o BadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableBadRequest struct {
	value *BadRequest
	isSet bool
}

func (v NullableBadRequest) Get() *BadRequest {
	return v.value
}

func (v *NullableBadRequest) Set(val *BadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequest(val *BadRequest) *NullableBadRequest {
	return &NullableBadRequest{value: val, isSet: true}
}

func (v NullableBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


