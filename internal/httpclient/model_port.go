/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kratos

import (
	"encoding/json"
)

// Port Port An open port on a container
type Port struct {
	// IP
	IP *string `json:"IP,omitempty"`
	// Port on the container
	PrivatePort int32 `json:"PrivatePort"`
	// Port exposed on the host
	PublicPort *int32 `json:"PublicPort,omitempty"`
	// type
	Type string `json:"Type"`
}

// NewPort instantiates a new Port object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPort(privatePort int32, type_ string) *Port {
	this := Port{}
	this.PrivatePort = privatePort
	this.Type = type_
	return &this
}

// NewPortWithDefaults instantiates a new Port object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortWithDefaults() *Port {
	this := Port{}
	return &this
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *Port) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *Port) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *Port) SetIP(v string) {
	o.IP = &v
}

// GetPrivatePort returns the PrivatePort field value
func (o *Port) GetPrivatePort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrivatePort
}

// GetPrivatePortOk returns a tuple with the PrivatePort field value
// and a boolean to check if the value has been set.
func (o *Port) GetPrivatePortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivatePort, true
}

// SetPrivatePort sets field value
func (o *Port) SetPrivatePort(v int32) {
	o.PrivatePort = v
}

// GetPublicPort returns the PublicPort field value if set, zero value otherwise.
func (o *Port) GetPublicPort() int32 {
	if o == nil || o.PublicPort == nil {
		var ret int32
		return ret
	}
	return *o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Port) GetPublicPortOk() (*int32, bool) {
	if o == nil || o.PublicPort == nil {
		return nil, false
	}
	return o.PublicPort, true
}

// HasPublicPort returns a boolean if a field has been set.
func (o *Port) HasPublicPort() bool {
	if o != nil && o.PublicPort != nil {
		return true
	}

	return false
}

// SetPublicPort gets a reference to the given int32 and assigns it to the PublicPort field.
func (o *Port) SetPublicPort(v int32) {
	o.PublicPort = &v
}

// GetType returns the Type field value
func (o *Port) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Port) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Port) SetType(v string) {
	o.Type = v
}

func (o Port) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IP != nil {
		toSerialize["IP"] = o.IP
	}
	if true {
		toSerialize["PrivatePort"] = o.PrivatePort
	}
	if o.PublicPort != nil {
		toSerialize["PublicPort"] = o.PublicPort
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePort struct {
	value *Port
	isSet bool
}

func (v NullablePort) Get() *Port {
	return v.value
}

func (v *NullablePort) Set(val *Port) {
	v.value = val
	v.isSet = true
}

func (v NullablePort) IsSet() bool {
	return v.isSet
}

func (v *NullablePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePort(val *Port) *NullablePort {
	return &NullablePort{value: val, isSet: true}
}

func (v NullablePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}