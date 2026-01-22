// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Attendee 
type Attendee struct {
    AttendeeBase
    // The proposedNewTime property
    proposedNewTime TimeSlotable
    // The status property
    status ResponseStatusable
}
// NewAttendee instantiates a new attendee and sets the default values.
func NewAttendee()(*Attendee) {
    m := &Attendee{
        AttendeeBase: *NewAttendeeBase(),
    }
    return m
}
// CreateAttendeeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttendeeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttendee(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Attendee) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AttendeeBase.GetFieldDeserializers()
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(TimeSlotable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResponseStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(ResponseStatusable))
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. The proposedNewTime property
func (m *Attendee) GetProposedNewTime()(TimeSlotable) {
    return m.proposedNewTime
}
// GetStatus gets the status property value. The status property
func (m *Attendee) GetStatus()(ResponseStatusable) {
    return m.status
}
// Serialize serializes information the current object
func (m *Attendee) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AttendeeBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProposedNewTime sets the proposedNewTime property value. The proposedNewTime property
func (m *Attendee) SetProposedNewTime(value TimeSlotable)() {
    m.proposedNewTime = value
}
// SetStatus sets the status property value. The status property
func (m *Attendee) SetStatus(value ResponseStatusable)() {
    m.status = value
}
// Attendeeable 
type Attendeeable interface {
    AttendeeBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProposedNewTime()(TimeSlotable)
    GetStatus()(ResponseStatusable)
    SetProposedNewTime(value TimeSlotable)()
    SetStatus(value ResponseStatusable)()
}
