// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package security

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for Authorization
func (this *Authorization) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Authorization
func (this *Authorization) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Group
func (this *Group) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Group
func (this *Group) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Rules
func (this *Rules) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Rules
func (this *Rules) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Match
func (this *Match) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Match
func (this *Match) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Address
func (this *Address) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Address
func (this *Address) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceAccountMatch
func (this *ServiceAccountMatch) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceAccountMatch
func (this *ServiceAccountMatch) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for StringMatch
func (this *StringMatch) MarshalJSON() ([]byte, error) {
	str, err := AuthorizationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StringMatch
func (this *StringMatch) UnmarshalJSON(b []byte) error {
	return AuthorizationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AuthorizationMarshaler   = &jsonpb.Marshaler{}
	AuthorizationUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
