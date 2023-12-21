package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ShiftAvailability 
type ShiftAvailability struct {
    // The recurrence property
    recurrence PatternedRecurrenceable
    // The time slot(s) preferred by the user.
    timeSlots []TimeRangeable
    // Specifies the time zone for the indicated time.
    timeZone *string
}
// NewShiftAvailability instantiates a new shiftAvailability and sets the default values.
func NewShiftAvailability()(*ShiftAvailability) {
    m := &ShiftAvailability{
    }
    return m
}
// CreateShiftAvailabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateShiftAvailabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShiftAvailability(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ShiftAvailability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["timeSlots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimeRangeable)
                }
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    return res
}
// GetRecurrence gets the recurrence property value. The recurrence property
func (m *ShiftAvailability) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetTimeSlots gets the timeSlots property value. The time slot(s) preferred by the user.
func (m *ShiftAvailability) GetTimeSlots()([]TimeRangeable) {
    return m.timeSlots
}
// GetTimeZone gets the timeZone property value. Specifies the time zone for the indicated time.
func (m *ShiftAvailability) GetTimeZone()(*string) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *ShiftAvailability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    if m.GetTimeSlots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecurrence sets the recurrence property value. The recurrence property
func (m *ShiftAvailability) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetTimeSlots sets the timeSlots property value. The time slot(s) preferred by the user.
func (m *ShiftAvailability) SetTimeSlots(value []TimeRangeable)() {
    m.timeSlots = value
}
// SetTimeZone sets the timeZone property value. Specifies the time zone for the indicated time.
func (m *ShiftAvailability) SetTimeZone(value *string)() {
    m.timeZone = value
}
// ShiftAvailabilityable 
type ShiftAvailabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecurrence()(PatternedRecurrenceable)
    GetTimeSlots()([]TimeRangeable)
    GetTimeZone()(*string)
    SetRecurrence(value PatternedRecurrenceable)()
    SetTimeSlots(value []TimeRangeable)()
    SetTimeZone(value *string)()
}
