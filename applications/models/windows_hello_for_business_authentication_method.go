// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsHelloForBusinessAuthenticationMethod 
type WindowsHelloForBusinessAuthenticationMethod struct {
    AuthenticationMethod
    // The date and time that this Windows Hello for Business key was registered.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The device property
    device Deviceable
    // The name of the device on which Windows Hello for Business is registered
    displayName *string
    // The keyStrength property
    keyStrength *AuthenticationMethodKeyStrength
}
// NewWindowsHelloForBusinessAuthenticationMethod instantiates a new windowsHelloForBusinessAuthenticationMethod and sets the default values.
func NewWindowsHelloForBusinessAuthenticationMethod()(*WindowsHelloForBusinessAuthenticationMethod) {
    m := &WindowsHelloForBusinessAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsHelloForBusinessAuthenticationMethod(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this Windows Hello for Business key was registered.
func (m *WindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDevice gets the device property value. The device property
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDevice()(Deviceable) {
    return m.device
}
// GetDisplayName gets the displayName property value. The name of the device on which Windows Hello for Business is registered
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsHelloForBusinessAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Deviceable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["keyStrength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodKeyStrength)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStrength(val.(*AuthenticationMethodKeyStrength))
        }
        return nil
    }
    return res
}
// GetKeyStrength gets the keyStrength property value. The keyStrength property
func (m *WindowsHelloForBusinessAuthenticationMethod) GetKeyStrength()(*AuthenticationMethodKeyStrength) {
    return m.keyStrength
}
// Serialize serializes information the current object
func (m *WindowsHelloForBusinessAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyStrength() != nil {
        cast := (*m.GetKeyStrength()).String()
        err = writer.WriteStringValue("keyStrength", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this Windows Hello for Business key was registered.
func (m *WindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDevice sets the device property value. The device property
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDevice(value Deviceable)() {
    m.device = value
}
// SetDisplayName sets the displayName property value. The name of the device on which Windows Hello for Business is registered
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetKeyStrength sets the keyStrength property value. The keyStrength property
func (m *WindowsHelloForBusinessAuthenticationMethod) SetKeyStrength(value *AuthenticationMethodKeyStrength)() {
    m.keyStrength = value
}
// WindowsHelloForBusinessAuthenticationMethodable 
type WindowsHelloForBusinessAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDevice()(Deviceable)
    GetDisplayName()(*string)
    GetKeyStrength()(*AuthenticationMethodKeyStrength)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDevice(value Deviceable)()
    SetDisplayName(value *string)()
    SetKeyStrength(value *AuthenticationMethodKeyStrength)()
}
