/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"encoding/json"
)

// checks if the ServiceSonarrSonarrIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSonarrSonarrIdGet200Response{}

// ServiceSonarrSonarrIdGet200Response struct for ServiceSonarrSonarrIdGet200Response
type ServiceSonarrSonarrIdGet200Response struct {
	Server *SonarrSettings `json:"server,omitempty"`
	Profiles *ServiceProfile `json:"profiles,omitempty"`
}

// NewServiceSonarrSonarrIdGet200Response instantiates a new ServiceSonarrSonarrIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSonarrSonarrIdGet200Response() *ServiceSonarrSonarrIdGet200Response {
	this := ServiceSonarrSonarrIdGet200Response{}
	return &this
}

// NewServiceSonarrSonarrIdGet200ResponseWithDefaults instantiates a new ServiceSonarrSonarrIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSonarrSonarrIdGet200ResponseWithDefaults() *ServiceSonarrSonarrIdGet200Response {
	this := ServiceSonarrSonarrIdGet200Response{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ServiceSonarrSonarrIdGet200Response) GetServer() SonarrSettings {
	if o == nil || IsNil(o.Server) {
		var ret SonarrSettings
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSonarrSonarrIdGet200Response) GetServerOk() (*SonarrSettings, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ServiceSonarrSonarrIdGet200Response) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given SonarrSettings and assigns it to the Server field.
func (o *ServiceSonarrSonarrIdGet200Response) SetServer(v SonarrSettings) {
	o.Server = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *ServiceSonarrSonarrIdGet200Response) GetProfiles() ServiceProfile {
	if o == nil || IsNil(o.Profiles) {
		var ret ServiceProfile
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSonarrSonarrIdGet200Response) GetProfilesOk() (*ServiceProfile, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ServiceSonarrSonarrIdGet200Response) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given ServiceProfile and assigns it to the Profiles field.
func (o *ServiceSonarrSonarrIdGet200Response) SetProfiles(v ServiceProfile) {
	o.Profiles = &v
}

func (o ServiceSonarrSonarrIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceSonarrSonarrIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return toSerialize, nil
}

type NullableServiceSonarrSonarrIdGet200Response struct {
	value *ServiceSonarrSonarrIdGet200Response
	isSet bool
}

func (v NullableServiceSonarrSonarrIdGet200Response) Get() *ServiceSonarrSonarrIdGet200Response {
	return v.value
}

func (v *NullableServiceSonarrSonarrIdGet200Response) Set(val *ServiceSonarrSonarrIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSonarrSonarrIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSonarrSonarrIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSonarrSonarrIdGet200Response(val *ServiceSonarrSonarrIdGet200Response) *NullableServiceSonarrSonarrIdGet200Response {
	return &NullableServiceSonarrSonarrIdGet200Response{value: val, isSet: true}
}

func (v NullableServiceSonarrSonarrIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSonarrSonarrIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


