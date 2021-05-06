/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DictExpression Used to create and directly specify the elements of a dictionary
type DictExpression struct {
	// Type of AST node
	Type *string `json:"type,omitempty"`
	// Elements of the dictionary
	Elements *[]DictItem `json:"elements,omitempty"`
}

// NewDictExpression instantiates a new DictExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictExpression() *DictExpression {
	this := DictExpression{}
	return &this
}

// NewDictExpressionWithDefaults instantiates a new DictExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictExpressionWithDefaults() *DictExpression {
	this := DictExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DictExpression) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictExpression) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DictExpression) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DictExpression) SetType(v string) {
	o.Type = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *DictExpression) GetElements() []DictItem {
	if o == nil || o.Elements == nil {
		var ret []DictItem
		return ret
	}
	return *o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictExpression) GetElementsOk() (*[]DictItem, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *DictExpression) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []DictItem and assigns it to the Elements field.
func (o *DictExpression) SetElements(v []DictItem) {
	o.Elements = &v
}

func (o DictExpression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableDictExpression struct {
	value *DictExpression
	isSet bool
}

func (v NullableDictExpression) Get() *DictExpression {
	return v.value
}

func (v *NullableDictExpression) Set(val *DictExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableDictExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableDictExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictExpression(val *DictExpression) *NullableDictExpression {
	return &NullableDictExpression{value: val, isSet: true}
}

func (v NullableDictExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDictExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}