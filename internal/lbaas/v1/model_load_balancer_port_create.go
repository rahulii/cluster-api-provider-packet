/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Load Balancer Management API

Load Balancer Management API is an API for managing load balancers.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the LoadBalancerPortCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPortCreate{}

// LoadBalancerPortCreate struct for LoadBalancerPortCreate
type LoadBalancerPortCreate struct {
	// A name for the port
	Name string `json:"name"`
	// Listing port
	Number int32 `json:"number"`
	// A list of pool IDs assigned to the port
	PoolIds              []string `json:"pool_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPortCreate LoadBalancerPortCreate

// NewLoadBalancerPortCreate instantiates a new LoadBalancerPortCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPortCreate(name string, number int32) *LoadBalancerPortCreate {
	this := LoadBalancerPortCreate{}
	this.Name = name
	this.Number = number
	return &this
}

// NewLoadBalancerPortCreateWithDefaults instantiates a new LoadBalancerPortCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPortCreateWithDefaults() *LoadBalancerPortCreate {
	this := LoadBalancerPortCreate{}
	return &this
}

// GetName returns the Name field value
func (o *LoadBalancerPortCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerPortCreate) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *LoadBalancerPortCreate) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortCreate) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *LoadBalancerPortCreate) SetNumber(v int32) {
	o.Number = v
}

// GetPoolIds returns the PoolIds field value if set, zero value otherwise.
func (o *LoadBalancerPortCreate) GetPoolIds() []string {
	if o == nil || IsNil(o.PoolIds) {
		var ret []string
		return ret
	}
	return o.PoolIds
}

// GetPoolIdsOk returns a tuple with the PoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortCreate) GetPoolIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PoolIds) {
		return nil, false
	}
	return o.PoolIds, true
}

// HasPoolIds returns a boolean if a field has been set.
func (o *LoadBalancerPortCreate) HasPoolIds() bool {
	if o != nil && !IsNil(o.PoolIds) {
		return true
	}

	return false
}

// SetPoolIds gets a reference to the given []string and assigns it to the PoolIds field.
func (o *LoadBalancerPortCreate) SetPoolIds(v []string) {
	o.PoolIds = v
}

func (o LoadBalancerPortCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPortCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["number"] = o.Number
	if !IsNil(o.PoolIds) {
		toSerialize["pool_ids"] = o.PoolIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPortCreate) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPortCreate := _LoadBalancerPortCreate{}

	err = json.Unmarshal(bytes, &varLoadBalancerPortCreate)

	if err != nil {
		return err
	}

	*o = LoadBalancerPortCreate(varLoadBalancerPortCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "pool_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPortCreate struct {
	value *LoadBalancerPortCreate
	isSet bool
}

func (v NullableLoadBalancerPortCreate) Get() *LoadBalancerPortCreate {
	return v.value
}

func (v *NullableLoadBalancerPortCreate) Set(val *LoadBalancerPortCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPortCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPortCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPortCreate(val *LoadBalancerPortCreate) *NullableLoadBalancerPortCreate {
	return &NullableLoadBalancerPortCreate{value: val, isSet: true}
}

func (v NullableLoadBalancerPortCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPortCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}