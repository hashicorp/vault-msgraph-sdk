// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttendeeBase 
type AttendeeBase struct {
    Recipient
    // The type property
    typeEscaped *AttendeeType
}
// NewAttendeeBase instantiates a new attendeeBase and sets the default values.
func NewAttendeeBase()(*AttendeeBase) {
    m := &AttendeeBase{
        Recipient: *NewRecipient(),
    }
    return m
}
// CreateAttendeeBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttendeeBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttendeeBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendeeBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Recipient.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttendeeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AttendeeType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
func (m *AttendeeBase) GetTypeEscaped()(*AttendeeType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AttendeeBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Recipient.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTypeEscaped sets the type property value. The type property
func (m *AttendeeBase) SetTypeEscaped(value *AttendeeType)() {
    m.typeEscaped = value
}
// AttendeeBaseable 
type AttendeeBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Recipientable
    GetTypeEscaped()(*AttendeeType)
    SetTypeEscaped(value *AttendeeType)()
}
