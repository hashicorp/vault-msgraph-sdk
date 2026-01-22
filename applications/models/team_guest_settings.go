// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamGuestSettings 
type TeamGuestSettings struct {
    // If set to true, guests can add and update channels.
    allowCreateUpdateChannels *bool
    // If set to true, guests can delete channels.
    allowDeleteChannels *bool
}
// NewTeamGuestSettings instantiates a new teamGuestSettings and sets the default values.
func NewTeamGuestSettings()(*TeamGuestSettings) {
    m := &TeamGuestSettings{
    }
    return m
}
// CreateTeamGuestSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamGuestSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamGuestSettings(), nil
}
// GetAllowCreateUpdateChannels gets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) GetAllowCreateUpdateChannels()(*bool) {
    return m.allowCreateUpdateChannels
}
// GetAllowDeleteChannels gets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) GetAllowDeleteChannels()(*bool) {
    return m.allowDeleteChannels
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamGuestSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowCreateUpdateChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateChannels(val)
        }
        return nil
    }
    res["allowDeleteChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteChannels(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamGuestSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCreateUpdateChannels", m.GetAllowCreateUpdateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteChannels", m.GetAllowDeleteChannels())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowCreateUpdateChannels sets the allowCreateUpdateChannels property value. If set to true, guests can add and update channels.
func (m *TeamGuestSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// SetAllowDeleteChannels sets the allowDeleteChannels property value. If set to true, guests can delete channels.
func (m *TeamGuestSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
// TeamGuestSettingsable 
type TeamGuestSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCreateUpdateChannels()(*bool)
    GetAllowDeleteChannels()(*bool)
    SetAllowCreateUpdateChannels(value *bool)()
    SetAllowDeleteChannels(value *bool)()
}
