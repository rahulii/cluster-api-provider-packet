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

// checks if the LoadBalancerPoolOriginCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPoolOriginCollection{}

// LoadBalancerPoolOriginCollection struct for LoadBalancerPoolOriginCollection
type LoadBalancerPoolOriginCollection struct {
	Origins              []LoadBalancerPoolOrigin `json:"origins"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPoolOriginCollection LoadBalancerPoolOriginCollection

// NewLoadBalancerPoolOriginCollection instantiates a new LoadBalancerPoolOriginCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPoolOriginCollection(origins []LoadBalancerPoolOrigin) *LoadBalancerPoolOriginCollection {
	this := LoadBalancerPoolOriginCollection{}
	this.Origins = origins
	return &this
}

// NewLoadBalancerPoolOriginCollectionWithDefaults instantiates a new LoadBalancerPoolOriginCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolOriginCollectionWithDefaults() *LoadBalancerPoolOriginCollection {
	this := LoadBalancerPoolOriginCollection{}
	return &this
}

// GetOrigins returns the Origins field value
func (o *LoadBalancerPoolOriginCollection) GetOrigins() []LoadBalancerPoolOrigin {
	if o == nil {
		var ret []LoadBalancerPoolOrigin
		return ret
	}

	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOriginCollection) GetOriginsOk() ([]LoadBalancerPoolOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origins, true
}

// SetOrigins sets field value
func (o *LoadBalancerPoolOriginCollection) SetOrigins(v []LoadBalancerPoolOrigin) {
	o.Origins = v
}

func (o LoadBalancerPoolOriginCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPoolOriginCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origins"] = o.Origins

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPoolOriginCollection) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPoolOriginCollection := _LoadBalancerPoolOriginCollection{}

	err = json.Unmarshal(bytes, &varLoadBalancerPoolOriginCollection)

	if err != nil {
		return err
	}

	*o = LoadBalancerPoolOriginCollection(varLoadBalancerPoolOriginCollection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "origins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPoolOriginCollection struct {
	value *LoadBalancerPoolOriginCollection
	isSet bool
}

func (v NullableLoadBalancerPoolOriginCollection) Get() *LoadBalancerPoolOriginCollection {
	return v.value
}

func (v *NullableLoadBalancerPoolOriginCollection) Set(val *LoadBalancerPoolOriginCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPoolOriginCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPoolOriginCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPoolOriginCollection(val *LoadBalancerPoolOriginCollection) *NullableLoadBalancerPoolOriginCollection {
	return &NullableLoadBalancerPoolOriginCollection{value: val, isSet: true}
}

func (v NullableLoadBalancerPoolOriginCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPoolOriginCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
