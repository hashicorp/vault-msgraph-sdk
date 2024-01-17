// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhoneAuthenticationMethod 
type PhoneAuthenticationMethod struct {
    AuthenticationMethod
    // The phone number to text or call for authentication. Phone numbers use the format +{country code} {number}x{extension}, with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating or updating if they don't match the required format.
    phoneNumber *string
    // The phoneType property
    phoneType *AuthenticationPhoneType
    // The smsSignInState property
    smsSignInState *AuthenticationMethodSignInState
}
// NewPhoneAuthenticationMethod instantiates a new phoneAuthenticationMethod and sets the default values.
func NewPhoneAuthenticationMethod()(*PhoneAuthenticationMethod) {
    m := &PhoneAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// CreatePhoneAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhoneAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhoneAuthenticationMethod(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhoneAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["phoneType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationPhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneType(val.(*AuthenticationPhoneType))
        }
        return nil
    }
    res["smsSignInState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodSignInState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsSignInState(val.(*AuthenticationMethodSignInState))
        }
        return nil
    }
    return res
}
// GetPhoneNumber gets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format +{country code} {number}x{extension}, with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating or updating if they don't match the required format.
func (m *PhoneAuthenticationMethod) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPhoneType gets the phoneType property value. The phoneType property
func (m *PhoneAuthenticationMethod) GetPhoneType()(*AuthenticationPhoneType) {
    return m.phoneType
}
// GetSmsSignInState gets the smsSignInState property value. The smsSignInState property
func (m *PhoneAuthenticationMethod) GetSmsSignInState()(*AuthenticationMethodSignInState) {
    return m.smsSignInState
}
// Serialize serializes information the current object
func (m *PhoneAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPhoneType() != nil {
        cast := (*m.GetPhoneType()).String()
        err = writer.WriteStringValue("phoneType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSmsSignInState() != nil {
        cast := (*m.GetSmsSignInState()).String()
        err = writer.WriteStringValue("smsSignInState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPhoneNumber sets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format +{country code} {number}x{extension}, with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating or updating if they don't match the required format.
func (m *PhoneAuthenticationMethod) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPhoneType sets the phoneType property value. The phoneType property
func (m *PhoneAuthenticationMethod) SetPhoneType(value *AuthenticationPhoneType)() {
    m.phoneType = value
}
// SetSmsSignInState sets the smsSignInState property value. The smsSignInState property
func (m *PhoneAuthenticationMethod) SetSmsSignInState(value *AuthenticationMethodSignInState)() {
    m.smsSignInState = value
}
// PhoneAuthenticationMethodable 
type PhoneAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPhoneNumber()(*string)
    GetPhoneType()(*AuthenticationPhoneType)
    GetSmsSignInState()(*AuthenticationMethodSignInState)
    SetPhoneNumber(value *string)()
    SetPhoneType(value *AuthenticationPhoneType)()
    SetSmsSignInState(value *AuthenticationMethodSignInState)()
}
