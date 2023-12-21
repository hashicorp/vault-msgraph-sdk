package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecurrenceRange 
type RecurrenceRange struct {
    // The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of times to repeat the event. Required and must be positive if type is numbered.
    numberOfOccurrences *int32
    // Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
    recurrenceTimeZone *string
    // The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
    startDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The type property
    typeEscaped *RecurrenceRangeType
}
// NewRecurrenceRange instantiates a new recurrenceRange and sets the default values.
func NewRecurrenceRange()(*RecurrenceRange) {
    m := &RecurrenceRange{
    }
    return m
}
// CreateRecurrenceRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecurrenceRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecurrenceRange(), nil
}
// GetEndDate gets the endDate property value. The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
func (m *RecurrenceRange) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecurrenceRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["numberOfOccurrences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfOccurrences(val)
        }
        return nil
    }
    res["recurrenceTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceTimeZone(val)
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrenceRangeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RecurrenceRangeType))
        }
        return nil
    }
    return res
}
// GetNumberOfOccurrences gets the numberOfOccurrences property value. The number of times to repeat the event. Required and must be positive if type is numbered.
func (m *RecurrenceRange) GetNumberOfOccurrences()(*int32) {
    return m.numberOfOccurrences
}
// GetRecurrenceTimeZone gets the recurrenceTimeZone property value. Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
func (m *RecurrenceRange) GetRecurrenceTimeZone()(*string) {
    return m.recurrenceTimeZone
}
// GetStartDate gets the startDate property value. The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
func (m *RecurrenceRange) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startDate
}
// GetTypeEscaped gets the type property value. The type property
func (m *RecurrenceRange) GetTypeEscaped()(*RecurrenceRangeType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RecurrenceRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numberOfOccurrences", m.GetNumberOfOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceTimeZone", m.GetRecurrenceTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDate sets the endDate property value. The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
func (m *RecurrenceRange) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetNumberOfOccurrences sets the numberOfOccurrences property value. The number of times to repeat the event. Required and must be positive if type is numbered.
func (m *RecurrenceRange) SetNumberOfOccurrences(value *int32)() {
    m.numberOfOccurrences = value
}
// SetRecurrenceTimeZone sets the recurrenceTimeZone property value. Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
func (m *RecurrenceRange) SetRecurrenceTimeZone(value *string)() {
    m.recurrenceTimeZone = value
}
// SetStartDate sets the startDate property value. The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
func (m *RecurrenceRange) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startDate = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RecurrenceRange) SetTypeEscaped(value *RecurrenceRangeType)() {
    m.typeEscaped = value
}
// RecurrenceRangeable 
type RecurrenceRangeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetNumberOfOccurrences()(*int32)
    GetRecurrenceTimeZone()(*string)
    GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTypeEscaped()(*RecurrenceRangeType)
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetNumberOfOccurrences(value *int32)()
    SetRecurrenceTimeZone(value *string)()
    SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTypeEscaped(value *RecurrenceRangeType)()
}
