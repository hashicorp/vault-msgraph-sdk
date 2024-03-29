// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentitySet 
type IdentitySet struct {
    // The application property
    application Identityable
    // The device property
    device Identityable
    // The user property
    user Identityable
}
// NewIdentitySet instantiates a new identitySet and sets the default values.
func NewIdentitySet()(*IdentitySet) {
    m := &IdentitySet{
    }
    return m
}
// CreateIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentitySet(), nil
}
// GetApplication gets the application property value. The application property
func (m *IdentitySet) GetApplication()(Identityable) {
    return m.application
}
// GetDevice gets the device property value. The device property
func (m *IdentitySet) GetDevice()(Identityable) {
    return m.device
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Identityable))
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Identityable))
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. The user property
func (m *IdentitySet) GetUser()(Identityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *IdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *IdentitySet) SetApplication(value Identityable)() {
    m.application = value
}
// SetDevice sets the device property value. The device property
func (m *IdentitySet) SetDevice(value Identityable)() {
    m.device = value
}
// SetUser sets the user property value. The user property
func (m *IdentitySet) SetUser(value Identityable)() {
    m.user = value
}
// IdentitySetable 
type IdentitySetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(Identityable)
    GetDevice()(Identityable)
    GetUser()(Identityable)
    SetApplication(value Identityable)()
    SetDevice(value Identityable)()
    SetUser(value Identityable)()
}
