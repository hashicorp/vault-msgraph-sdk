// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeZoneBase 
type TimeZoneBase struct {
    // The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
    name *string
}
// NewTimeZoneBase instantiates a new timeZoneBase and sets the default values.
func NewTimeZoneBase()(*TimeZoneBase) {
    m := &TimeZoneBase{
    }
    return m
}
// CreateTimeZoneBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeZoneBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeZoneBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeZoneBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *TimeZoneBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) SetName(value *string)() {
    m.name = value
}
// TimeZoneBaseable 
type TimeZoneBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    SetName(value *string)()
}
