// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeSlot 
type TimeSlot struct {
    // The end property
    end DateTimeTimeZoneable
    // The start property
    start DateTimeTimeZoneable
}
// NewTimeSlot instantiates a new timeSlot and sets the default values.
func NewTimeSlot()(*TimeSlot) {
    m := &TimeSlot{
    }
    return m
}
// CreateTimeSlotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeSlotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeSlot(), nil
}
// GetEnd gets the end property value. The end property
func (m *TimeSlot) GetEnd()(DateTimeTimeZoneable) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeSlot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetStart gets the start property value. The start property
func (m *TimeSlot) GetStart()(DateTimeTimeZoneable) {
    return m.start
}
// Serialize serializes information the current object
func (m *TimeSlot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnd sets the end property value. The end property
func (m *TimeSlot) SetEnd(value DateTimeTimeZoneable)() {
    m.end = value
}
// SetStart sets the start property value. The start property
func (m *TimeSlot) SetStart(value DateTimeTimeZoneable)() {
    m.start = value
}
// TimeSlotable 
type TimeSlotable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnd()(DateTimeTimeZoneable)
    GetStart()(DateTimeTimeZoneable)
    SetEnd(value DateTimeTimeZoneable)()
    SetStart(value DateTimeTimeZoneable)()
}
