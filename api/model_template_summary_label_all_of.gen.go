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

// TemplateSummaryLabelAllOf struct for TemplateSummaryLabelAllOf
type TemplateSummaryLabelAllOf struct {
	Id         uint64                              `json:"id" yaml:"id"`
	OrgID      *uint64                             `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	Name       string                              `json:"name" yaml:"name"`
	Properties TemplateSummaryLabelAllOfProperties `json:"properties" yaml:"properties"`
}

// NewTemplateSummaryLabelAllOf instantiates a new TemplateSummaryLabelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryLabelAllOf(id uint64, name string, properties TemplateSummaryLabelAllOfProperties) *TemplateSummaryLabelAllOf {
	this := TemplateSummaryLabelAllOf{}
	this.Id = id
	this.Name = name
	this.Properties = properties
	return &this
}

// NewTemplateSummaryLabelAllOfWithDefaults instantiates a new TemplateSummaryLabelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryLabelAllOfWithDefaults() *TemplateSummaryLabelAllOf {
	this := TemplateSummaryLabelAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateSummaryLabelAllOf) GetId() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOf) GetIdOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryLabelAllOf) SetId(v uint64) {
	o.Id = v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *TemplateSummaryLabelAllOf) GetOrgID() uint64 {
	if o == nil || o.OrgID == nil {
		var ret uint64
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOf) GetOrgIDOk() (*uint64, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *TemplateSummaryLabelAllOf) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given int64 and assigns it to the OrgID field.
func (o *TemplateSummaryLabelAllOf) SetOrgID(v uint64) {
	o.OrgID = &v
}

// GetName returns the Name field value
func (o *TemplateSummaryLabelAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryLabelAllOf) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value
func (o *TemplateSummaryLabelAllOf) GetProperties() TemplateSummaryLabelAllOfProperties {
	if o == nil {
		var ret TemplateSummaryLabelAllOfProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOf) GetPropertiesOk() (*TemplateSummaryLabelAllOfProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *TemplateSummaryLabelAllOf) SetProperties(v TemplateSummaryLabelAllOfProperties) {
	o.Properties = v
}

func (o TemplateSummaryLabelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryLabelAllOf struct {
	value *TemplateSummaryLabelAllOf
	isSet bool
}

func (v NullableTemplateSummaryLabelAllOf) Get() *TemplateSummaryLabelAllOf {
	return v.value
}

func (v *NullableTemplateSummaryLabelAllOf) Set(val *TemplateSummaryLabelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryLabelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryLabelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryLabelAllOf(val *TemplateSummaryLabelAllOf) *NullableTemplateSummaryLabelAllOf {
	return &NullableTemplateSummaryLabelAllOf{value: val, isSet: true}
}

func (v NullableTemplateSummaryLabelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryLabelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
