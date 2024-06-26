/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"time"
)

// checks if the AllOfPrimitiveTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllOfPrimitiveTypes{}

// AllOfPrimitiveTypes struct for AllOfPrimitiveTypes
type AllOfPrimitiveTypes struct {
	Test *time.Time `json:"test,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllOfPrimitiveTypes AllOfPrimitiveTypes

// NewAllOfPrimitiveTypes instantiates a new AllOfPrimitiveTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllOfPrimitiveTypes() *AllOfPrimitiveTypes {
	this := AllOfPrimitiveTypes{}
	return &this
}

// NewAllOfPrimitiveTypesWithDefaults instantiates a new AllOfPrimitiveTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllOfPrimitiveTypesWithDefaults() *AllOfPrimitiveTypes {
	this := AllOfPrimitiveTypes{}
	return &this
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *AllOfPrimitiveTypes) GetTest() time.Time {
	if o == nil || IsNil(o.Test) {
		var ret time.Time
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOfPrimitiveTypes) GetTestOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *AllOfPrimitiveTypes) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given time.Time and assigns it to the Test field.
func (o *AllOfPrimitiveTypes) SetTest(v time.Time) {
	o.Test = &v
}

func (o AllOfPrimitiveTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllOfPrimitiveTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllOfPrimitiveTypes) UnmarshalJSON(bytes []byte) (err error) {
	varAllOfPrimitiveTypes := _AllOfPrimitiveTypes{}

	err = json.Unmarshal(bytes, &varAllOfPrimitiveTypes)

	if err != nil {
		return err
	}

	*o = AllOfPrimitiveTypes(varAllOfPrimitiveTypes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllOfPrimitiveTypes struct {
	value *AllOfPrimitiveTypes
	isSet bool
}

func (v NullableAllOfPrimitiveTypes) Get() *AllOfPrimitiveTypes {
	return v.value
}

func (v *NullableAllOfPrimitiveTypes) Set(val *AllOfPrimitiveTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAllOfPrimitiveTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAllOfPrimitiveTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllOfPrimitiveTypes(val *AllOfPrimitiveTypes) *NullableAllOfPrimitiveTypes {
	return &NullableAllOfPrimitiveTypes{value: val, isSet: true}
}

func (v NullableAllOfPrimitiveTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllOfPrimitiveTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


