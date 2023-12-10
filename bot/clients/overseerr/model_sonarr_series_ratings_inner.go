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

// checks if the SonarrSeriesRatingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarrSeriesRatingsInner{}

// SonarrSeriesRatingsInner struct for SonarrSeriesRatingsInner
type SonarrSeriesRatingsInner struct {
	Votes *float32 `json:"votes,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewSonarrSeriesRatingsInner instantiates a new SonarrSeriesRatingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarrSeriesRatingsInner() *SonarrSeriesRatingsInner {
	this := SonarrSeriesRatingsInner{}
	return &this
}

// NewSonarrSeriesRatingsInnerWithDefaults instantiates a new SonarrSeriesRatingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarrSeriesRatingsInnerWithDefaults() *SonarrSeriesRatingsInner {
	this := SonarrSeriesRatingsInner{}
	return &this
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *SonarrSeriesRatingsInner) GetVotes() float32 {
	if o == nil || IsNil(o.Votes) {
		var ret float32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesRatingsInner) GetVotesOk() (*float32, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *SonarrSeriesRatingsInner) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given float32 and assigns it to the Votes field.
func (o *SonarrSeriesRatingsInner) SetVotes(v float32) {
	o.Votes = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SonarrSeriesRatingsInner) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesRatingsInner) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SonarrSeriesRatingsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *SonarrSeriesRatingsInner) SetValue(v float32) {
	o.Value = &v
}

func (o SonarrSeriesRatingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarrSeriesRatingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSonarrSeriesRatingsInner struct {
	value *SonarrSeriesRatingsInner
	isSet bool
}

func (v NullableSonarrSeriesRatingsInner) Get() *SonarrSeriesRatingsInner {
	return v.value
}

func (v *NullableSonarrSeriesRatingsInner) Set(val *SonarrSeriesRatingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarrSeriesRatingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarrSeriesRatingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarrSeriesRatingsInner(val *SonarrSeriesRatingsInner) *NullableSonarrSeriesRatingsInner {
	return &NullableSonarrSeriesRatingsInner{value: val, isSet: true}
}

func (v NullableSonarrSeriesRatingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarrSeriesRatingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


