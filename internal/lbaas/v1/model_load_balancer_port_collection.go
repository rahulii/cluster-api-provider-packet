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

// checks if the LoadBalancerPortCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPortCollection{}

// LoadBalancerPortCollection struct for LoadBalancerPortCollection
type LoadBalancerPortCollection struct {
	Ports                []LoadBalancerPort `json:"ports"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPortCollection LoadBalancerPortCollection

// NewLoadBalancerPortCollection instantiates a new LoadBalancerPortCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPortCollection(ports []LoadBalancerPort) *LoadBalancerPortCollection {
	this := LoadBalancerPortCollection{}
	this.Ports = ports
	return &this
}

// NewLoadBalancerPortCollectionWithDefaults instantiates a new LoadBalancerPortCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPortCollectionWithDefaults() *LoadBalancerPortCollection {
	this := LoadBalancerPortCollection{}
	return &this
}

// GetPorts returns the Ports field value
func (o *LoadBalancerPortCollection) GetPorts() []LoadBalancerPort {
	if o == nil {
		var ret []LoadBalancerPort
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortCollection) GetPortsOk() ([]LoadBalancerPort, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *LoadBalancerPortCollection) SetPorts(v []LoadBalancerPort) {
	o.Ports = v
}

func (o LoadBalancerPortCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPortCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ports"] = o.Ports

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPortCollection) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPortCollection := _LoadBalancerPortCollection{}

	err = json.Unmarshal(bytes, &varLoadBalancerPortCollection)

	if err != nil {
		return err
	}

	*o = LoadBalancerPortCollection(varLoadBalancerPortCollection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPortCollection struct {
	value *LoadBalancerPortCollection
	isSet bool
}

func (v NullableLoadBalancerPortCollection) Get() *LoadBalancerPortCollection {
	return v.value
}

func (v *NullableLoadBalancerPortCollection) Set(val *LoadBalancerPortCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPortCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPortCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPortCollection(val *LoadBalancerPortCollection) *NullableLoadBalancerPortCollection {
	return &NullableLoadBalancerPortCollection{value: val, isSet: true}
}

func (v NullableLoadBalancerPortCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPortCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
