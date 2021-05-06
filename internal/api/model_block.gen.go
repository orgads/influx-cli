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

// Block A set of statements
type Block struct {
	// Type of AST node
	Type *string `json:"type,omitempty"`
	// Block body
	Body *[]Statement `json:"body,omitempty"`
}

// NewBlock instantiates a new Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlock() *Block {
	this := Block{}
	return &this
}

// NewBlockWithDefaults instantiates a new Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWithDefaults() *Block {
	this := Block{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Block) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Block) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Block) SetType(v string) {
	o.Type = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Block) GetBody() []Statement {
	if o == nil || o.Body == nil {
		var ret []Statement
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetBodyOk() (*[]Statement, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Block) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []Statement and assigns it to the Body field.
func (o *Block) SetBody(v []Statement) {
	o.Body = &v
}

func (o Block) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableBlock struct {
	value *Block
	isSet bool
}

func (v NullableBlock) Get() *Block {
	return v.value
}

func (v *NullableBlock) Set(val *Block) {
	v.value = val
	v.isSet = true
}

func (v NullableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlock(val *Block) *NullableBlock {
	return &NullableBlock{value: val, isSet: true}
}

func (v NullableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}